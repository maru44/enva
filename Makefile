.PHONY: build/cli compression defrost test

# BIN
BIN_DIR:=bin/
BIN_NAME:=enva
ifeq (${CLI_API_URL}, http://localhost:8080)
  BIN_NAME=enva_dev
endif

# FILE
ENVA_MAIN_FILE:=./enva/main.go

# GOOS
GOOS_LINUX:=linux
GOOS_WINDOWS:=windows
GOOS_DARWIN:=darwin
# GOARCH
GOARCH_AMD:=amd64

# test dirs
# test dirs
TEST_USECASE:=./service/api/internal/usecase
TEST_REPOSITORY:=./service/api/internal/interface/database
TEST_CONTROLLER:=./service/api/internal/interface/controllers
TEST_INFRA:=./service/api/internal/infra

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

test:
	@go test $(TEST_REPOSITORY) $(TEST_CONTROLLER) $(TEST_USECASE) $(TEST_INFRA)
