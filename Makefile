.PHONY: all get test clean build cover

GO ?= go
BIN_NAME=parser
VERSION=$(shell git describe --tags --always)

all: clean build

get:
	${GO} get

build: get
	# ${GO} get -u gopkg.in/yaml.v2;
	${GO} build -ldflags "-X main.version=${VERSION}" -o ${BIN_NAME} cmd/parser/main.go;

clean:
	@rm -rf ${BIN_NAME} *.out build

test: get
	${GO} test -race -v

cover:
	${GO} test -cover && \
	${GO} test -coverprofile=coverage.out  && \
	${GO} tool cover -html=coverage.out
