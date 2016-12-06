set PKGNAME=github.com/oneumyvakin/debindices
set LOCALPATH=%~dp0

go fmt %PKGNAME%
goimports.exe -w .
staticcheck.exe %PKGNAME%

set GOOS=linux
set GOARCH=amd64
go build -o deb.%GOARCH% %PKGNAME%

set GOOS=windows
set GOARCH=amd64
go build -o deb.exe %PKGNAME%