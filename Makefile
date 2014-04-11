#
# Makefile for gosigma
#

PROJECT := github.com/Altoros/gosigma
PROJECT_DIR := $(shell go list -e -f '{{.Dir}}' $(PROJECT))

ifeq ($(shell uname -p | sed -r 's/.*(x86|armel|armhf).*/golang/'), golang)
	GO_C := golang
	INSTALL_FLAGS := 
else
	GO_C := gccgo-4.9  gccgo-go
	INSTALL_FLAGS := -gccgoflags=-static-libgo
endif

default: build

# Start of GOPATH-dependent targets. Some targets only make sense -
# and will only work - when this tree is found on the GOPATH.
ifeq ($(CURDIR),$(PROJECT_DIR))

build:
	go build $(PROJECT)/...

check test:
	go test $(PROJECT)/...

install:
	go install $(INSTALL_FLAGS) -v $(PROJECT)/...

clean:
	go clean $(PROJECT)/...

else # --------------------------------

build check test install clean:
	$(error Cannot $@; $(CURDIR) is not on GOPATH)

endif
# End of GOPATH-dependent targets.

# Reformat source files.
format:
	gofmt -w -l .

.PHONY: build check test install clean
.PHONY: format


