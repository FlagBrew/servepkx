# Go parameters
GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOGET = $(GOCMD) get

SOURCE = src
BINARY_NAME = servepkx
BINARY_WINDOWS = $(BINARY_NAME).exe
BINARY_LINUX = $(BINARY_NAME)_linux
BINARY_MAC = $(BINARY_NAME)_mac

all: deps build

build:
	@cd $(SOURCE) && GOOS=windows GOARCH=amd64 $(GOBUILD) -o ../$(BINARY_WINDOWS)
	@cd $(SOURCE) && GOOS=linux GOARCH=amd64 $(GOBUILD) -o ../$(BINARY_LINUX)
	@cd $(SOURCE) && GOOS=darwin GOARCH=amd64 $(GOBUILD) -o ../$(BINARY_MAC)

clean:
	@$(GOCLEAN)
	@rm -f $(BINARY_WINDOWS)
	@rm -f $(BINARY_LINUX)
	@rm -f $(BINARY_MAC)
	@echo Clean...

deps:
	@$(GOGET) github.com/schollz/progressbar