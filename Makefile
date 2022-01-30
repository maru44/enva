.PHONY: build/cli \
	compression \
	defrost \
	test \
	touch/tar \
	json/version \
	backup \
	container/image \
	migrate/up \
	migrate/fix \
	migrate/drop \
	migrate/version

# ADMIN CLI GO FILE
ADMIN:=./service/admin/internal/main.go

# BIN
BIN_DIR:=bin/
BIN_NAME:=enva
TAR_DIR:=./service/front/public/enva/
ifeq (${ENVIRONMENT}, development)
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

VERSION:=$(shell jq .version ./enva/commands/version.json)

# test dirs
# test dirs
TEST_USECASE:=./service/api/internal/usecase
TEST_REPOSITORY:=./service/api/internal/interface/database
TEST_CONTROLLER:=./service/api/internal/interface/controllers
TEST_INFRA:=./service/api/internal/infra

build/cli:
	@echo 'Start ${GOOS_DARWIN}'
	@GOOS=${GOOS_DARWIN} GOARCH=${GOARCH_AMD} go build -o ${BIN_DIR}enva/${GOOS_DARWIN}/${GOARCH_AMD}/${BIN_NAME} ${ENVA_MAIN_FILE}
	@echo 'Finish ${GOOS_DARWIN}!!'

	@echo 'Start ${GOOS_LINUX}'
	@GOOS=${GOOS_LINUX} GOARCH=${GOARCH_AMD} go build -o ${BIN_DIR}enva/${GOOS_LINUX}/${GOARCH_AMD}/${BIN_NAME} ${ENVA_MAIN_FILE}
	@GOOS=${GOOS_LINUX} GOARCH=${GOARCH_386} go build -o ${BIN_DIR}enva/${GOOS_LINUX}/${GOARCH_386}/${BIN_NAME} ${ENVA_MAIN_FILE}
	@echo 'Finish ${GOOS_LINUX}!!'

	@echo 'Start ${GOOS_WINDOWS}'
	@GOOS=${GOOS_WINDOWS} GOARCH=${GOARCH_AMD} go build -o ${BIN_DIR}enva/${GOOS_WINDOWS}/${GOARCH_AMD}/${BIN_NAME} ${ENVA_MAIN_FILE}
	@GOOS=${GOOS_WINDOWS} GOARCH=${GOARCH_386} go build -o ${BIN_DIR}enva/${GOOS_WINDOWS}/${GOARCH_386}/${BIN_NAME} ${ENVA_MAIN_FILE}
	@echo 'Finish ${GOOS_WINDOWS}!!'

touch/tar:
	@touch ${TAR_DIR}enva_${VERSION}_${GOOS_DARWIN}_${GOARCH_AMD}.tar.gz
	@touch ${TAR_DIR}enva_${VERSION}_${GOOS_LINUX}_${GOARCH_AMD}.tar.gz
	@touch ${TAR_DIR}enva_${VERSION}_${GOOS_LINUX}_${GOARCH_386}.tar.gz
	@touch ${TAR_DIR}enva_${VERSION}_${GOOS_WINDOWS}_${GOARCH_AMD}.tar.gz
	@touch ${TAR_DIR}enva_${VERSION}_${GOOS_WINDOWS}_${GOARCH_386}.tar.gz

compress:
	@echo 'Compression ${GOOS_DARWIN}'
	@tar -C ${BIN_DIR}enva/${GOOS_DARWIN}/${GOARCH_AMD}/ -cvzf ${TAR_DIR}enva_${VERSION}_${GOOS_DARWIN}_${GOARCH_AMD}.tar.gz ${BIN_NAME}
	@go run ${ADMIN} tar/json ${VERSION} ${GOOS_DARWIN} ${GOARCH_AMD}

	@echo 'Compression ${GOOS_LINUX}'
	@tar -C ${BIN_DIR}enva/${GOOS_LINUX}/${GOARCH_AMD}/ -cvzf ${TAR_DIR}enva_${VERSION}_${GOOS_LINUX}_${GOARCH_AMD}.tar.gz ${BIN_NAME}
	@tar -C ${BIN_DIR}enva/${GOOS_LINUX}/${GOARCH_386}/ -cvzf ${TAR_DIR}enva_${VERSION}_${GOOS_LINUX}_${GOARCH_386}.tar.gz ${BIN_NAME}
	@go run ${ADMIN} tar/json ${VERSION} ${GOOS_LINUX} ${GOARCH_AMD}
	@go run ${ADMIN} tar/json ${VERSION} ${GOOS_LINUX} ${GOARCH_386}

	@echo 'Compression ${GOOS_WINDOWS}'
	@tar -C ${BIN_DIR}enva/${GOOS_WINDOWS}/${GOARCH_AMD}/ -cvzf ${TAR_DIR}enva_${VERSION}_${GOOS_WINDOWS}_${GOARCH_AMD}.tar.gz ${BIN_NAME}
	@tar -C ${BIN_DIR}enva/${GOOS_WINDOWS}/${GOARCH_386}/ -cvzf ${TAR_DIR}enva_${VERSION}_${GOOS_WINDOWS}_${GOARCH_386}.tar.gz ${BIN_NAME}
	@go run ${ADMIN} ${VERSION} ${GOOS_WINDOWS} ${GOARCH_AMD}
	@go run ${ADMIN} tar/json ${VERSION} ${GOOS_WINDOWS} ${GOARCH_386}

defrost:
	@tar -xvzf ${BIN_DIR}enva/enva_darwin.tar.gz 

test:
	@go test $(TEST_REPOSITORY) $(TEST_CONTROLLER) $(TEST_USECASE) $(TEST_INFRA)

explain/json:
	@go run ${ADMIN} explain/json

privacy/json:
	@go run ${ADMIN} privacy/json

backup:
	@go run ${ADMIN} backup

container/image:
	@docker-compose -f docker-compose.go.image.yaml build
	@docker tag ${ECR_REPOSITORY_API} ${ECR_REGISTRY_API}/${ECR_REPOSITORY_API}

container/push:
	@docker push ${ECR_REGISTRY_API}/${ECR_REPOSITORY_API}:latest

migrate/up:
	@echo 'migration (up)'
	@go run ${ADMIN} migrate up

migrate/fix:
	@echo 'migration (fix)'
	@go run ${ADMIN} migrate fix

migrate/drop:
	@echo 'migration (drop)'
	@go run ${ADMIN} migrate drop

migrate/version:
	@go run ${ADMIN} migrate version
