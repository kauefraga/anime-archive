BINARY_FOLDER=bin
BINARY_NAME=anime-archive


# Build with optimizations (-ldflags=-w and cgo disabled)
release:
	GOOS=darwin go build -ldflags=-w -o ${BINARY_FOLDER}/darwin/${BINARY_NAME} ./cmd/main.go
	GOOS=linux go build -ldflags=-w -o ${BINARY_FOLDER}/linux/${BINARY_NAME} ./cmd/main.go
	GOOS=windows go build -ldflags=-w -o ${BINARY_FOLDER}/windows/${BINARY_NAME}.exe ./cmd/main.go
