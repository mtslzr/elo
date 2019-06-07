GOCMD=go
GOBUILD=${GOCMD} build
GOLINT=golangci-lint
GOMOD=${GOCMD} mod
GOTEST=${GOCMD} test

PACKAGE=elo

all: lint test build

build:
	${GOBUILD} -o ${PACKAGE} -v

db:
	cat ${SQL} | ${DBCMD} ${DB}

deps: tidy vend

dump:
	${DBCMD} ${DB} .dump > ${SQL}

lint:
	${GOLINT} run

prep: deps dump

test:
	${GOTEST} -v ./...

run: build
	./${PACKAGE}

tidy:
	${GOMOD} tidy -v

vend:
	${GOMOD} vendor -v