.PHONY: build/cli compression defrost print

# BIN
BIN_DIR:=bin/
BIN_NAME:=enva
ifeq (${CLI_API_URL}, http://localhost:8080)
  BIN_NAME=enva_div
endif

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
	@GOOS=${GOOS_DARWIN} GOARCH=${GOARCH_AMD} go build -o ${BIN_DIR}enva/${GOOS_DARWIN}/${BIN_NAME} ${ENVA_MAIN_FILE}
	@echo 'Finish ${GOOS_DARWIN}!!'

	@echo 'Start ${GOOS_LINUX}'
	@GOOS=${GOOS_LINUX} GOARCH=${GOARCH_AMD} go build -o ${BIN_DIR}enva/${GOOS_LINUX}/${BIN_NAME} ${ENVA_MAIN_FILE}
	@echo 'Finish ${GOOS_LINUX}!!'

	@echo 'Start ${GOOS_WINDOWS}'
	@GOOS=${GOOS_WINDOWS} GOARCH=${GOARCH_AMD} go build -o ${BIN_DIR}enva/${GOOS_WINDOWS}/${BIN_NAME} ${ENVA_MAIN_FILE}
	@echo 'Finish ${GOOS_WINDOWS}!!'

compress:
	@echo 'Compression darwin'
	@cd ${BIN_DIR}
	@tar -cvzf ${BIN_DIR}enva/enva_${GOOS_DARWIN}.tar.gz ./enva/${GOOS_DARWIN}/${BIN_NAME}
	@cd -

defrost:
	@tar -xvzf ${BIN_DIR}enva/enva_darwin.tar.gz 
