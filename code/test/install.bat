#echo off

setloacal

if exist install.bat goto ok
echo install.bat must be run from its folder
goto end

:OK

set OLDGOPATH=%GOPATH%
set GOPATH=%~dp0

gofmt -w src

go install test

:end
echo finished