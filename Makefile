ifeq ($(OS),Windows_NT)
	BINARY_NAME=mssql-ddlextractor.exe
else
	BINARY_NAME=mssql-ddlextractor
endif

build:
	@echo Compiling...
	go mod tidy
	go build -o bin/${BINARY_NAME} cmd/main.go
	@echo Done!