# Synopsis: Build sources
task build {
    Set-Location -Path 'src'
    Start-Process -FilePath 'sb' -ArgumentList 'gen' -NoNewWindow -Wait
    Start-Process -FilePath 'sb' -ArgumentList 'build' -NoNewWindow -Wait
}

# Synopsis: Build samples
task samples {
    Set-Location -Path 'src\samples'
    Start-Process -FilePath 'sb' -ArgumentList 'gen helloworld' -NoNewWindow -Wait
    Start-Process -FilePath 'sb' -ArgumentList 'build helloworld' -NoNewWindow -Wait
}

# Synopsis: Install sb application
task install {
    Set-Location -Path 'src\sb'
    Copy-Item -Path 'sb.exe' -Destination '..\..\bin\'
    Start-Process -FilePath 'go' -ArgumentList 'install' -NoNewWindow -Wait
}

# Synopsis: Run tests
task test {
    Set-Location -Path 'src\tests\cmd'
    Start-Process -FilePath 'go' -ArgumentList 'test' -NoNewWindow -Wait
}

task . build, test, samples