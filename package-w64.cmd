@echo off
setlocal
set pacf=jvnman_%~1_windows_64bit.zip
set GO111MODULE=on
echo package to %pacf%
pause

go test -v ./...
if ERRORLEVEL 1 echo fail go test && goto :end

go run main.go version
go build -o dist/jvnman.exe -ldflags "-X github.com/spiegel-im-spiegel/jvnman/facade.Version=%~1" .
if ERRORLEVEL 1 echo fail go build && goto :end

pushd dist
jvnman.exe version

copy /v /b ..\LICENSE .
copy /v /b ..\README.md .
7z.exe a -tzip %pacf% jvnman.exe LICENSE README.md
if ERRORLEVEL 1 echo fail packaging %pacf% && goto :end

gpg --detach-sign -u spiegel.im.spiegel@gmail.com -o %pacf%.sig %pacf%
if ERRORLEVEL 1 echo fail signing %pacf% && goto :end

gpg -d %pacf%.sig

:end
popd
endlocal 
exit /b 0
