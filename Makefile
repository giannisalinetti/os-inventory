GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test ./...
BINFILE=os-inventory

build:
	$(GOTEST) && \
	$(GOBUILD) -o $(BINFILE) -v

clean:
	$(GOCLEAN)
	rm -f $(BINFILE)

test:
	$(GOTEST) -v

install:
	mv $(BINFILE) /usr/local/bin
