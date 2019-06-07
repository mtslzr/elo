GOCMD=go
GOBUILD=${GOCMD} build
GOLINT=golangci-lint
GOMOD=${GOCMD} mod
GOTEST=${GOCMD} test

PACKAGE=elo

all: lint test build

build:
	${GOBUILD} -o ${PACKAGE} -v

deps: tidy vend

lint:
	${GOLINT} run

test:
	${GOTEST} -v ./...

run: build
	./${PACKAGE}

tidy:
	${GOMOD} tidy -v

vend:
	${GOMOD} vendor -v