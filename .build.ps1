
# Synopsis: Generate sources
task gen gen-sb, gen-sgo

# Synopsis: Generate sb sources
task gen-sb {
    Set-Location -Path 'src'
    $Status = Start-Process -FilePath 'sb' -ArgumentList 'gen sb' -NoNewWindow -PassThru -Wait
    Assert($Status.ExitCode -eq 0) 'The "gen sb" command failed'
}

# Synopsis: Generate sgo sources
task gen-sgo {
    Set-Location -Path 'src'
    $Status = Start-Process -FilePath 'sb' -ArgumentList 'gen sgo' -NoNewWindow -PassThru -Wait
    Assert($Status.ExitCode -eq 0) 'The "gen sgo" command failed'
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
task gbuild gbuild-sb, gbuild-sgo

# Synopsis: Generate & build sb application
task gbuild-sb gen-sb, build-sb

# Synopsis: Generate & build sgo plugin
task gbuild-sgo gen-sgo, build-sgo

# Synopsis: Build samples
task samples {
    Set-Location -Path 'src\samples'
    $Status = Start-Process -FilePath 'sb' -ArgumentList 'gen helloworld' -NoNewWindow -PassThru -Wait
    Assert($Status.ExitCode -eq 0) 'The "gen helloworld" command failed'
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
task ginstall gbuild, install

# Synopsis: Generate, build & install sb application
task ginstall-sb gbuild-sb, install-sb

# Synopsis: Generate, build & install sgo plugim
task ginstall-sgo gbuild-sgo, install-sgo

# Synopsis: Run tests
task test {
    Set-Location -Path 'src\tests\cmd'
    $Status = Start-Process -FilePath 'go' -ArgumentList 'test' -NoNewWindow -PassThru -Wait
    Assert($Status.ExitCode -eq 0) 'The test command failed'
}

task . gbuild, test, samples