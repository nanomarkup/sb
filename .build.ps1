
# Synopsis: Generate sources
task gen -Before build {
    Set-Location -Path 'src'
    $Status = Start-Process -FilePath 'sb' -ArgumentList 'gen' -NoNewWindow -PassThru -Wait
    Assert($Status.ExitCode -eq 0) 'The gen command failed'
}

# Synopsis: Build sources
task build {
    Set-Location -Path 'src'
    $Status = Start-Process -FilePath 'sb' -ArgumentList 'build' -NoNewWindow -PassThru -Wait 
    Assert($Status.ExitCode -eq 0) 'The build command failed'
}

# Synopsis: Build samples
task samples {
    Set-Location -Path 'src\samples'
    $Status = Start-Process -FilePath 'sb' -ArgumentList 'gen helloworld' -NoNewWindow -PassThru -Wait
    Assert($Status.ExitCode -eq 0) 'The gen command failed'
    $Status = Start-Process -FilePath 'sb' -ArgumentList 'build helloworld' -NoNewWindow -PassThru -Wait
    Assert($Status.ExitCode -eq 0) 'The build command failed'
}

# Synopsis: Install sb application
task install {
    Set-Location -Path 'src\sb'
    Copy-Item -Path 'sb.exe' -Destination '..\..\bin\'
    $Status = Start-Process -FilePath 'go' -ArgumentList 'install' -NoNewWindow -PassThru -Wait
    Assert($Status.ExitCode -eq 0) 'The install command failed'
}

# Synopsis: Run tests
task test {
    Set-Location -Path 'src\tests\cmd'
    $Status = Start-Process -FilePath 'go' -ArgumentList 'test' -NoNewWindow -PassThru -Wait
    Assert($Status.ExitCode -eq 0) 'The test command failed'
}

task . build, test, samples