.PHONY: build/cli

# BIN
BIN_DIR:=bin/
BIN_NAME:=main

# FILE
ENVA_MAIN_FILE:=./enva/main.go

# GOOS
GOOS_LINUX:=linux
GOOS_WINDOWS:=windows
GOOS_DARWIN:=darwin
# GOARCH
GOARCH_AMD:=amd64

build/cli:
	@echo 'Start ${GOOS_DARWIN}'
	@GOOS=${GOOS_DARWIN} GOARCH=${GOARCH_AMD} go build -o ${BIN_DIR}enva/${GOOS_DARWIN}/enva/${BIN_NAME} ${ENVA_MAIN_FILE}
	@echo 'Finish ${GOOS_DARWIN}!!'

	@echo 'Start ${GOOS_LINUX}'
	@GOOS=${GOOS_LINUX} GOARCH=${GOARCH_AMD} go build -o ${BIN_DIR}enva/${GOOS_LINUX}/enva/${BIN_NAME} ${ENVA_MAIN_FILE}
	@echo 'Finish ${GOOS_LINUX}!!'

	@echo 'Start ${GOOS_WINDOWS}'
	@GOOS=${GOOS_WINDOWS} GOARCH=${GOARCH_AMD} go build -o ${BIN_DIR}enva/${GOOS_WINDOWS}/enva/${BIN_NAME} ${ENVA_MAIN_FILE}
	@echo 'Finish ${GOOS_WINDOWS}!!'
