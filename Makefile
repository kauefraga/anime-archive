BINARY_FOLDER=bin
BINARY_NAME=anime-archive

build:
	GOARCH=amd64 GOOS=darwin go build -o ${BINARY_FOLDER}/${BINARY_NAME}-darwin ./src/main.go
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_FOLDER}/${BINARY_NAME}-linux ./src/main.go
	GOARCH=amd64 GOOS=windows go build -o ${BINARY_FOLDER}/${BINARY_NAME}.exe ./src/main.go

release:
	GOARCH=amd64 GOOS=darwin go build -ldflags=-w -o ${BINARY_FOLDER}/${BINARY_NAME}-darwin ./src/main.go
	GOARCH=amd64 GOOS=linux go build -ldflags=-w -o ${BINARY_FOLDER}/${BINARY_NAME}-linux ./src/main.go
	GOARCH=amd64 GOOS=windows go build -ldflags=-w -o ${BINARY_FOLDER}/${BINARY_NAME}.exe ./src/main.go

clean:
	go clean
	rm ./${BINARY_FOLDER}/${BINARY_NAME}-darwin
	rm ./${BINARY_FOLDER}/${BINARY_NAME}-linux
	rm ./${BINARY_FOLDER}/${BINARY_NAME}.exe
