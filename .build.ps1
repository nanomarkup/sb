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
task code {
    $Status = Start-Process -FilePath 'sb' -ArgumentList 'code' -NoNewWindow -PassThru -Wait
    Assert($Status.ExitCode -eq 0) 'The "code" command failed'
}

# Synopsis: Build sources
task build {
    $Status = Start-Process -FilePath 'sb' -ArgumentList 'build' -NoNewWindow -PassThru -Wait 
    Assert($Status.ExitCode -eq 0) 'The "build" command failed'
}

# Synopsis: Generate & build sources
task cbuild code, build

# Synopsis: Remove generated files
task clean {
    $AppPath = './sb/sb'
    if ($PSVersionTable.Platform -ne 'Unix') {
        $AppPath += '.exe'
    }
    if (Test-Path -Path $AppPath) {
        Remove-Item -Path $AppPath
    }
}

# Synopsis: Build samples
task build-samples {
    Set-Location -Path 'samples'
    $Status = Start-Process -FilePath 'sb' -ArgumentList 'code helloworld' -NoNewWindow -PassThru -Wait
    Assert($Status.ExitCode -eq 0) 'The "code helloworld" command failed'
    $Status = Start-Process -FilePath 'sb' -ArgumentList 'build helloworld' -NoNewWindow -PassThru -Wait
    Assert($Status.ExitCode -eq 0) 'The "build helloworld" command failed'
}

# Synopsis: Clean samples
task clean-samples {
    Set-Location -Path 'samples'
    $Status = Start-Process -FilePath 'sb' -ArgumentList 'clean helloworld' -NoNewWindow -PassThru -Wait
    Assert($Status.ExitCode -eq 0) 'The "clean helloworld" command failed'
}

# Synopsis: Remove all generated files
task clean-all clean, clean-samples

# Synopsis: Install application
task install {
    $AppName = 'sb'
    if ($PSVersionTable.Platform -eq 'Unix') {
        $GoPath = "${Env:HOME}/go"
    } else {        
        $GoPath = "${Env:GOPATH}".TrimEnd(';')
        $AppName += '.exe'
    }
    Set-Location -Path 'sb'
    Copy-Item -Path $AppName -Destination '../bin/'    
    Copy-Item -Path $AppName -Destination "$GoPath/bin/"
}

# Synopsis: Generate, build & install application
task cinstall cbuild, install

# Synopsis: Run tests
task test {
    Set-Location -Path 'cmd'
    $Status = Start-Process -FilePath 'go' -ArgumentList 'test' -NoNewWindow -PassThru -Wait
    Assert($Status.ExitCode -eq 0) 'The test command failed'
}

# Synopsis: Generate documentation
task doc {
    GenDoc -PackageName '.'
    GenDoc -PackageName 'cmd'
    GenDoc -PackageName 'plugins'
}

task . cbuild, build-samples, test, clean-all, doc