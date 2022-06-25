
# Synopsis: Generate sources
task gen -Before build {
    Set-Location -Path 'src'
    $Status = Start-Process -FilePath 'sb' -ArgumentList 'gen sb' -NoNewWindow -PassThru -Wait
    Assert($Status.ExitCode -eq 0) 'The "gen sb" command failed'
    $Status = Start-Process -FilePath 'sb' -ArgumentList 'gen sgo' -NoNewWindow -PassThru -Wait
    Assert($Status.ExitCode -eq 0) 'The "gen sgo" command failed'
}

# Synopsis: Build sources
task build {
    Set-Location -Path 'src'
    $Status = Start-Process -FilePath 'sb' -ArgumentList 'build sb' -NoNewWindow -PassThru -Wait 
    Assert($Status.ExitCode -eq 0) 'The "build sb" command failed'
    $Status = Start-Process -FilePath 'sb' -ArgumentList 'build sgo' -NoNewWindow -PassThru -Wait 
    Assert($Status.ExitCode -eq 0) 'The "build sgo" command failed'
}

# Synopsis: Build samples
task samples {
    Set-Location -Path 'src\samples'
    $Status = Start-Process -FilePath 'sb' -ArgumentList 'gen helloworld' -NoNewWindow -PassThru -Wait
    Assert($Status.ExitCode -eq 0) 'The "gen helloworld" command failed'
    $Status = Start-Process -FilePath 'sb' -ArgumentList 'build helloworld' -NoNewWindow -PassThru -Wait
    Assert($Status.ExitCode -eq 0) 'The "build helloworld" command failed'
}

# Synopsis: Install sb application
task install {
    $GoPath = "${Env:GOPATH}".TrimEnd(';')
    Set-Location -Path 'src\sb'
    Copy-Item -Path 'sb.exe' -Destination '..\..\bin\'    
    Copy-Item -Path 'sb.exe' -Destination "$GoPath\bin\"
    Set-Location -Path '..\sgo'
    Copy-Item -Path 'sgo.exe' -Destination '..\..\bin\'
    Copy-Item -Path 'sgo.exe' -Destination "$GoPath\bin\"
}

# Synopsis: Run tests
task test {
    Set-Location -Path 'src\tests\cmd'
    $Status = Start-Process -FilePath 'go' -ArgumentList 'test' -NoNewWindow -PassThru -Wait
    Assert($Status.ExitCode -eq 0) 'The test command failed'
}

task . build, test, samples