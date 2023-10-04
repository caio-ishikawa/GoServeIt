APP_NAME := gsi
SOURCE_FILE := main.go
INSTALL_DIR := /usr/local/bin

build:
	go build -o $(INSTALL_DIR)/$(APP_NAME) $(SOURCE_FILE)

clean:
	rm -rf $(INSTALL_DIR)/$(APP_NAME)

uninstall:
	rm -rf $(INSTALL_DIR)/$(APP_NAME)

.PHONY: build clean
