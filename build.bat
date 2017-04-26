@echo off

set GOOS=windows
set GOARCH=386

set BUILD_ARGS=-ldflags "-H windowsgui"
set BINDATA_ARGS=-nomemcopy

go-bindata %BINDATA_ARGS% 123456.jpg
go build -o YLP2.exe %BUILD_ARGS%

go-bindata %BINDATA_ARGS% -prefix "NC" NC
go build -o YLPNC.exe %BUILD_ARGS%
