OS = linux
GO_BUILD = go build
GOFLAGS = GOOS=${OS} CGO_ENABLED=0

build:
	${GOFLAGS} ${GO_BUILD} -tags netgo -a -v -installsuffix cgo -o bin/organize cmd/main.go
.PHONY: go-build

windows:
	GOOS=windows CGO_ENABLED=0 -tags netgo -a -v -installsuffix cgo -o bin/organize cmd/main.go
.PHONY: go-windows

vendor:
	go mod vendor
.PHONY: vendor

all: vendor build
.PHONY: all

default: all
.PHONY: default

