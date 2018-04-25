GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
BINFILE=os-inventory

build:
	$(GOBUILD) -o $(BINFILE) -v

clean:
	$(GOCLEAN)
	rm -f $(BINFILE)

install:
	mv $(BINFILE) /usr/local/bin
