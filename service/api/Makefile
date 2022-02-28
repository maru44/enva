.PHONY: run build exe test

BIN_DIR:=../../bin/
BIN_NAME:=main
CMD_DIR:=cmd/

# root under cmd
SERVERROOT:=server

# test dirs
TEST_USECASE:=./internal/usecase
TEST_REPOSITORY:=./internal/interface/database
TEST_CONTROLLER:=./internal/interface/controllers
TEST_INFRA:=./internal/infra

# args for build or run
t:=$(SERVERROOT)

run:
	@echo "localhost:8080 is started"
	@go run $(CMD_DIR)$(SERVERROOT)/main.go

build:
	@go build -o $(BIN_DIR)${t}/$(BIN_NAME) $(CMD_DIR)${t}/main.go

exe:
	@./$(BIN_DIR)$(SERVERROOT)/$(BIN_NAME) &

test:
	@go test $(TEST_REPOSITORY) $(TEST_CONTROLLER) $(TEST_USECASE) $(TEST_INFRA)

gen:
	go generate ./...
