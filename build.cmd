@echo off
pushd .\src

rem Build a SmartBuilder application
sb gen
if %errorlevel%==1 goto error
sb build
if %errorlevel%==1 goto error
echo Built

rem Test the application
pushd .\tests\cmd
go test
if %errorlevel%==1 goto error
echo Tested
popd

rem Copy the application to the bin folder and install it
pushd .\sb
copy sb.exe ..\..\bin\ /y >NUL
if %errorlevel%==1 goto error
echo Delivered
go install
if %errorlevel%==1 goto error
echo Installed
popd

popd
goto success

:error
echo Failed
pause

:success
