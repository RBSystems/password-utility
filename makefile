INSTALL_DIR := $(GOPATH)/bin
BIN := pipass

all: 
	go build -o $(INSTALL_DIR)/$(BIN)



