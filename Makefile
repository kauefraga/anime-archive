BINARY_FOLDER=bin
BINARY_NAME=anime-archive


# Build with optimizations (-ldflags=-w)
release:
	GOOS=darwin go build -ldflags=-w -o ${BINARY_FOLDER}/darwin/${BINARY_NAME} ./src/main.go
	GOOS=linux go build -ldflags=-w -o ${BINARY_FOLDER}/linux/${BINARY_NAME} ./src/main.go
	GOOS=windows go build -ldflags=-w -o ${BINARY_FOLDER}/windows/${BINARY_NAME}.exe ./src/main.go
