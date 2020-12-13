set GOOS=windows
: set GOARCH=386
: set AR=i686-w64-mingw32-ar
: set CC=i686-w64-mingw32-gcc
: set CXX=i686-w64-mingw32-g++
: set GOARCH=amd64
: set AR=x86_64-w64-mingw32-ar
: set CC=x86_64-w64-mingw32-gcc
: set CXX=x86_64-w64-mingw32-g++
: set CGO_ENABLED=1
: @set PATH=%CYGWIN64DIR%\bin;%PATH%

set TARGET=sample

: go fmt %TARGET%
: go fmt %TARGET%/config
: go fmt %TARGET%/controller
: go fmt %TARGET%/model
: go fmt %TARGET%/startup

: golangci-lint run . config controller model startup

: windres -o %TARGET%.syso %TARGET%.rc

: go build -o build/%TARGET%.exe -race -ldflags "-H windowsgui -w -s" %TARGET%
go build -o build/%TARGET%.exe -ldflags "-w -s" %TARGET%

@pause