.PHONY: build/cli compression defrost test blank/tar

# BIN
BIN_DIR:=bin/
BIN_NAME:=enva
TAR_DIR:=./tar/
ifeq (${CLI_API_URL}, http://localhost:8080)
  BIN_NAME=enva_dev
  TAR_DIR=${BIN_DIR}enva/
endif

# FILE
ENVA_MAIN_FILE:=./enva/main.go

# GOOS
GOOS_LINUX:=linux
GOOS_WINDOWS:=windows
GOOS_DARWIN:=darwin
# GOARCH
GOARCH_AMD:=amd64
GOARCH_386:=386

# test dirs
# test dirs
TEST_USECASE:=./service/api/internal/usecase
TEST_REPOSITORY:=./service/api/internal/interface/database
TEST_CONTROLLER:=./service/api/internal/interface/controllers
TEST_INFRA:=./service/api/internal/infra

num:=0

build/cli:
	@echo 'Start ${GOOS_DARWIN}'
	@GOOS=${GOOS_DARWIN} GOARCH=${GOARCH_AMD} go build -o ${BIN_DIR}enva/${GOOS_DARWIN}_${GOARCH_AMD}/${BIN_NAME} ${ENVA_MAIN_FILE}
	@GOOS=${GOOS_DARWIN} GOARCH=${GOARCH_386} go build -o ${BIN_DIR}enva/${GOOS_DARWIN}_${GOARCH_386}/${BIN_NAME} ${ENVA_MAIN_FILE}
	@echo 'Finish ${GOOS_DARWIN}!!'

	@echo 'Start ${GOOS_LINUX}'
	@GOOS=${GOOS_LINUX} GOARCH=${GOARCH_AMD} go build -o ${BIN_DIR}enva/${GOOS_LINUX}_${GOARCH_AMD}/${BIN_NAME} ${ENVA_MAIN_FILE}
	@GOOS=${GOOS_LINUX} GOARCH=${GOARCH_386} go build -o ${BIN_DIR}enva/${GOOS_LINUX}_${GOARCH_386}/${BIN_NAME} ${ENVA_MAIN_FILE}
	@echo 'Finish ${GOOS_LINUX}!!'

	@echo 'Start ${GOOS_WINDOWS}'
	@GOOS=${GOOS_WINDOWS} GOARCH=${GOARCH_AMD} go build -o ${BIN_DIR}enva/${GOOS_WINDOWS}_${GOARCH_AMD}/${BIN_NAME} ${ENVA_MAIN_FILE}
	@GOOS=${GOOS_WINDOWS} GOARCH=${GOARCH_386} go build -o ${BIN_DIR}enva/${GOOS_WINDOWS}_${GOARCH_386}/${BIN_NAME} ${ENVA_MAIN_FILE}
	@echo 'Finish ${GOOS_WINDOWS}!!'

blank/tar:
	@echo ${num}

compress:
	@echo 'Compression ${GOOS_DARWIN}'
	@tar -C ${BIN_DIR}enva/${GOOS_DARWIN}/${GOARCH_AMD}/ -cvzf ${TAR_DIR}enva_${GOOS_DARWIN}_${GOARCH_AMD}.tar.gz ${BIN_NAME}
	@tar -C ${BIN_DIR}enva/${GOOS_DARWIN}/${GOARCH_386}/ -cvzf ${TAR_DIR}enva_${GOOS_DARWIN}_${GOARCH_386}.tar.gz ${BIN_NAME}
	@echo 'Compression ${GOOS_LINUX}'
	@tar -C ${BIN_DIR}enva/${GOOS_LINUX}/${GOARCH_AMD}/ -cvzf ${TAR_DIR}enva_${GOOS_LINUX}_${GOARCH_AMD}.tar.gz ${BIN_NAME}
	@tar -C ${BIN_DIR}enva/${GOOS_LINUX}/${GOARCH_386}/ -cvzf ${TAR_DIR}enva_${GOOS_LINUX}_${GOARCH_386}.tar.gz ${BIN_NAME}
	@echo 'Compression ${GOOS_WINDOWS}'
	@tar -C ${BIN_DIR}enva/${GOOS_WINDOWS}/${GOARCH_AMD}/ -cvzf ${TAR_DIR}enva_${GOOS_WINDOWS}_${GOARCH_AMD}.tar.gz ${BIN_NAME}
	@tar -C ${BIN_DIR}enva/${GOOS_WINDOWS}/${GOARCH_386}/ -cvzf ${TAR_DIR}enva_${GOOS_WINDOWS}_${GOARCH_386}.tar.gz ${BIN_NAME}

defrost:
	@tar -xvzf ${BIN_DIR}enva/enva_darwin.tar.gz 

test:
	@go test $(TEST_REPOSITORY) $(TEST_CONTROLLER) $(TEST_USECASE) $(TEST_INFRA)
