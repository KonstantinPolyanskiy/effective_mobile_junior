.PHONY: test build run clean
GOCMD = go
GOBUILD = $(GOCMD) build
GOTEST = $(GOCMD) test
GOBIN = ./bin
BINARY_NAME = effective_mobile_junior
MAIN_PACKAGE = ./cmd/app

all: test build

test:
	$(GOTEST) ./...

build:
	$(GOBUILD) -o $(GOBIN)/$(BINARY_NAME) $(MAIN_PACKAGE)

run: build
	$(GOBIN)/$(BINARY_NAME)

clean:
	rm -f $(GOBIN)/$(BINARY_NAME)
