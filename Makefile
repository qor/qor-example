OSNAME=$(shell uname)

GO=$(shell which go)

CUR_TIME=$(shell date '+%Y-%m-%d_%H:%M:%S')
# Program version
VERSION=$(shell cat VERSION)

# Binary name for bintray
BIN_NAME=$(shell basename $(abspath ./))

# Project name for bintray
PROJECT_NAME=$(shell basename $(abspath ./))
PROJECT_DIR=$(shell pwd)

# Project url used for builds
# examples: github.com, bitbucket.org
REPO_HOST_URL=github.com.org

# Grab the current commit
GIT_COMMIT="$(shell git rev-parse HEAD)"

# Check if there are uncommited changes
GIT_DIRTY="$(shell test -n "`git status --porcelain`" && echo "+CHANGES" || true)"

# Add the godep path to the GOPATH
#GOPATH=$(shell godep path):$(shell echo $$GOPATH)

default: help

help:
	@echo "..............................................................."
	@echo "Project: $(PROJECT_NAME) | current dir: $(PROJECT_DIR)"
	@echo "version: $(VERSION) GIT_DIRTY: $(GIT_DIRTY)\n"
	@#echo "Autocomplete exec -> PROG=$(BIN_NAME) source ./autocomplete/bash_autocomplete\n"
	@echo "make init    - Load godep"
	@echo "make save    - Save project libs"
	@echo "make install - Install packages"
	@echo "make clean   - Clean .orig, .log files"
	@echo "make run     - Run project debug mode"
	@echo "make seed    - Run project seeds"
	@echo "make build   - Build for current OS project"
	@echo "make release - Build release project"
	@#echo "make docs"   - Project documentation
	@echo "...............................................................\n"

init:
	@go get github.com/tools/godep

save:
	godep save

install:
	@go get -v -u
	@go get -v -u github.com/gin-gonic/gin
	@go get -v -u github.com/codegangsta/cli
	@go get -v -u github.com/azumads/faker

release:
	@echo "building release ${BIN_NAME} ${VERSION}"
	@echo "GOPATH=${GOPATH}"
	godep get && \
	 CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w' -o bin/${BIN_NAME} main.go

clean:
	@test ! -e ./${BIN_NAME} || rm ./${BIN_NAME}
	@git gc --prune=0 --aggressive
	@find . -name "*.orig" -type f -delete
	@find . -name "*.log" -type f -delete

seed:
	@echo "...............................................................\n"
	@echo $(PROJECT_NAME) seed
	@echo ...............................................................
	@MACARON_ENV=development go run db/seeds/main.go

run:
	@echo "...............................................................\n"
	@echo $(PROJECT_NAME)
	@echo documentation API open in browser:
	@echo	"	 http://localhost:7000/\n"
	@echo ...............................................................
	@MACARON_ENV=development go run main.go

test:
	@go test -v ./...
	@#API_PATH=$(PROJECT_DIR) ginkgo -v -r

push:
	@git add -A
	@git ci -am "new release v$(VERSION) COMMIT: $(GIT_COMMIT)"
	@git push

build: clean
	@echo "Building ${BIN_NAME} ${VERSION}"
	@CGO_ENABLED=0 go build -a -tags netgo -ldflags '-w -X main.BuildTime=${CUR_TIME} -X main.Version=${VERSION} -X main.GitHash=${GIT_COMMIT}' -o $(BIN_NAME) main.go

docs:
	godoc -http=:6060 -index

