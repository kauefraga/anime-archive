BINARY_FOLDER=bin
BINARY_NAME=anime-archive

# WINDOWS_SUFIX=.exe
LINUX_SUFIX=linux # .bin
DARWIN_SUFIX=darwin


install:
	go get github.com/adrg/xdg
	go get github.com/spf13/cobra
	go get github.com/jedib0t/go-pretty/v6
	go get gorm.io/driver/sqlite
	go get -u gorm.io/gorm

# Build with optimizations (-ldflags=-w)
release:
	GOARCH=amd64 GOOS=darwin go build -ldflags=-w -o ${BINARY_FOLDER}/${BINARY_NAME}-${DARWIN_SUFIX} ./src/main.go
	GOARCH=amd64 GOOS=linux go build -ldflags=-w -o ${BINARY_FOLDER}/${BINARY_NAME}-${LINUX_SUFIX} ./src/main.go
	GOARCH=amd64 GOOS=windows go build -ldflags=-w -o ${BINARY_FOLDER}/${BINARY_NAME}.exe ./src/main.go

# Build with optimizations (-ldflags=-w) for x86
release-x86:
	GOARCH=386 GOOS=linux go build -ldflags=-w -o ${BINARY_FOLDER}/${BINARY_NAME}-${LINUX_SUFIX}-x86 ./src/main.go
	GOARCH=386 GOOS=windows go build -ldflags=-w -o ${BINARY_FOLDER}/${BINARY_NAME}-x86.exe ./src/main.go
