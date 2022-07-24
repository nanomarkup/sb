function GenDoc {
    param (
        [string]$PackageName
    )
    $CurrLocation = Get-Location
    Set-Location -Path $PackageName
    $Status = Start-Process -FilePath 'go' -ArgumentList 'doc -all' -RedirectStandardOutput 'readme.txt' -NoNewWindow -PassThru -Wait 
    Set-Location -Path $CurrLocation
    Assert($Status.ExitCode -eq 0) 'The "go doc" command failed'
}

# Synopsis: Generate sources
task code code-sb, code-sgo

# Synopsis: Generate sb sources
task code-sb {
    Set-Location -Path 'src'
    $Status = Start-Process -FilePath 'sb' -ArgumentList 'code sb' -NoNewWindow -PassThru -Wait
    Assert($Status.ExitCode -eq 0) 'The "code sb" command failed'
}

# Synopsis: Generate sgo sources
task code-sgo {
    Set-Location -Path 'src'
    $Status = Start-Process -FilePath 'sb' -ArgumentList 'code sgo' -NoNewWindow -PassThru -Wait
    Assert($Status.ExitCode -eq 0) 'The "code sgo" command failed'
}

# Synopsis: Build sources
task build build-sb, build-sgo

# Synopsis: Build sb application
task build-sb {
    Set-Location -Path 'src'
    $Status = Start-Process -FilePath 'sb' -ArgumentList 'build sb' -NoNewWindow -PassThru -Wait 
    Assert($Status.ExitCode -eq 0) 'The "build sb" command failed'
}

# Synopsis: Build sgo plugin
task build-sgo {
    Set-Location -Path 'src'
    $Status = Start-Process -FilePath 'sb' -ArgumentList 'build sgo' -NoNewWindow -PassThru -Wait 
    Assert($Status.ExitCode -eq 0) 'The "build sgo" command failed'
}

# Synopsis: Generate & build sources
task cbuild cbuild-sb, cbuild-sgo

# Synopsis: Generate & build sb application
task cbuild-sb code-sb, build-sb

# Synopsis: Generate & build sgo plugin
task cbuild-sgo code-sgo, build-sgo

# Synopsis: Build samples
task samples {
    Set-Location -Path 'src\samples'
    $Status = Start-Process -FilePath 'sb' -ArgumentList 'code helloworld' -NoNewWindow -PassThru -Wait
    Assert($Status.ExitCode -eq 0) 'The "code helloworld" command failed'
    $Status = Start-Process -FilePath 'sb' -ArgumentList 'build helloworld' -NoNewWindow -PassThru -Wait
    Assert($Status.ExitCode -eq 0) 'The "build helloworld" command failed'
}

# Synopsis: Install applications & plugins
task install install-sb, install-sgo

# Synopsis: Install sb application
task install-sb {
    $GoPath = "${Env:GOPATH}".TrimEnd(';')
    Set-Location -Path 'src\sb'
    Copy-Item -Path 'sb.exe' -Destination '..\..\bin\'    
    Copy-Item -Path 'sb.exe' -Destination "$GoPath\bin\"
}

# Synopsis: Install sgo plugin
task install-sgo {
    $GoPath = "${Env:GOPATH}".TrimEnd(';')
    Set-Location -Path 'src\sgo'
    Copy-Item -Path 'sgo.exe' -Destination '..\..\bin\'
    Copy-Item -Path 'sgo.exe' -Destination "$GoPath\bin\"
}

# Synopsis: Generate, build & install applications 
task cinstall cbuild, install

# Synopsis: Generate, build & install sb application
task cinstall-sb cbuild-sb, install-sb

# Synopsis: Generate, build & install sgo plugim
task cinstall-sgo cbuild-sgo, install-sgo

# Synopsis: Run tests
task test {
    Set-Location -Path 'src\tests\cmd'
    $Status = Start-Process -FilePath 'go' -ArgumentList 'test' -NoNewWindow -PassThru -Wait
    Assert($Status.ExitCode -eq 0) 'The test command failed'
}

# Synopsis: Generate documentation
task doc {
    Set-Location -Path 'src'
    GenDoc -PackageName 'app'
    GenDoc -PackageName 'cmd'
    GenDoc -PackageName 'golang'
    GenDoc -PackageName 'helper\hashicorp\hclog'
    GenDoc -PackageName 'plugins'
    GenDoc -PackageName 'plugins\sgo'
    GenDoc -PackageName 'smodule'
}

task . cbuild, test, samples, doc