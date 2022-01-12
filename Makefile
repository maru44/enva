.PHONY: build/cli

BIN_DIR:=bin/
BIN_NAME:=main

build/cli:
	@go build -o ${BIN_DIR}enva/${BIN_NAME} ./enva/main.go
