DEPCMD=dep
DEPENSURE=$(DEPCMD) ensure

GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

BINARY_NAME=gitpls
DIST_PATH=$(CURDIR)/dist

.PHONY: all
all: test build

.PHONY: restore
restore:
	$(DEPENSURE)

.PHONY: build
build:
	mkdir -p $(DIST_PATH)
	$(GOBUILD) -o "$(DIST_PATH)/$(BINARY_NAME)" -v "github.com/nlowe/gitpls/cmd/gitpls"

.PHONY: test
test: 
	$(GOTEST) -v ./...

.PHONY: clean
clean: 
	$(GOCLEAN)
	rm -rf $(DIST_PATH)

.PHONY: run
run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)
