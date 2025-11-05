BIN_DIR := bin

SRC := cmd/api/main.go
OUT := $(BIN_DIR)/api

DI_PATH := internal/bootstrap/dependency_injection.go

prebuild:
	@go tool kessoku $(DI_PATH)

build: prebuild
	@go build -o $(OUT) $(SRC)

run: build
	@$(OUT)

clean:
	@rm -rf $(GO_BIN)

.PHONY: build run clean
