MODULES=activity l10n responder sorting audited location roles transition exchange media_library seo validations i18n qor serializable_meta worker inflection slug
OSNAME=$(shell uname)

GO=$(shell which go)

CUR_TIME=$(shell date '+%Y-%m-%d_%H:%M:%S')
# Program version
VERSION=$(shell cat VERSION)

# Binary name for bintray
BIN_NAME=$(shell basename $(abspath ./))
BIN_NAME_CLI=qor-cli

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
	@echo "make cli     - Build qor-cli"
	@echo "make build   - Build for current OS project"
	@echo "make release - Build release project"
	@echo "make docs    - Project documentation"
	@echo "...............................................................\n"

init:
	@go get github.com/tools/godep

save:
	@godep save

install:
	@go get -v -u github.com/constabulary/gb/...
	@go get -v -u github.com/kr/godep
	@go get -v -u github.com/gin-gonic/gin
	@go get -v -u github.com/codegangsta/cli
	@go get -v -u github.com/azumads/faker
	@go get -v -u github.com/jteeuwen/go-bindata/...
	@#go get -v -u
	@#go get -v -u

git:
	@for a in $(MODULES); do echo "-> $$a"; cd ../$$a && git pull; done

template:
	@mkdir -p ./dist/config
	@mkdir -p ./dist/app/views/qor
	@mkdir -p public/admin/assets
	@rm -R ./public/admin/assets
	@mkdir -p ./public/admin/assets/javascripts/vendors
	@cp -R ../qor/admin/views/* ./dist/app/views/qor/
	@cp -R ../activity/views/themes/activities/metas ./dist/app/views/qor/
	@cp -R ../i18n/exchange_actions/views/themes/i18n/actions ./dist/app/views/qor/
	@#cp ../i18n/views/themes/i18n/inline-edit-libs.tmpl.tmpl ./dist/app/views/qor/
	@#cp ../i18n/views/themes/i18n/index.tmpl ./dist/app/views/qor/
	@#cp ../l10n/views/themes/l10n/new.tmpl ./dist/app/views/qor/
	@#cp ../seo/views/themes/seo/edit.tmpl ./dist/app/views/qor/
	@#cp ../worker/views/themes/worker/edit.tmpl ./dist/app/views/qor/
	@#cp ../worker/views/themes/worker/new.tmpl ./dist/app/views/qor/
	@cp -R ../l10n/views/metas ./dist/app/views/qor/
	@cp -R ../l10n/views/themes/l10n/actions ./dist/app/views/qor/
	@cp -R ../location/views/metas ./dist/app/views/qor/
	@cp -R ../media_library/views/metas ./dist/app/views/qor/
	@cp -R ../seo/views/metas ./dist/app/views/qor/
	@cp -R ../seo/views/microdata ./dist/app/views/qor/
	@cp -R ../serializable_meta/views/metas ./dist/app/views/qor/
	@cp -R ../slug/views/metas ./dist/app/views/qor/
	@cp -R ../sorting/views/themes/sorting/actions ./dist/app/views/qor/
	@cp -R ../worker/views/themes/worker/actions ./dist/app/views/qor/
	@cp -R ./app/views/* ./dist/app/views/

assets:
	@cp ./config/database.yml ./dist/config/
	@cp -R ./dist/app/views/qor/assets ./public/admin/
	@cp ../qor/bower_components/jquery/dist/jquery.min.map ./public/admin/assets/javascripts/vendors/
	@cp -R ../activity/views/themes/activities/assets ./public/admin/
	@cp -R ../i18n/exchange_actions/views/assets ./public/admin/
	@cp -R ../i18n/views/themes/i18n/assets ./public/admin/
	@cp -R ../l10n/views/themes/l10n/assets ./public/admin/
	@cp -R ../location/views/themes/location/assets ./public/admin/
	@cp -R ../seo/views/themes/seo/assets ./public/admin/
	@cp -R ../seo/images ./public/admin/assets/
	@cp -R ../slug/views/themes/slug/assets ./public/admin/
	@cp -R ../sorting/views/themes/sorting/assets ./public/admin/
	@cp -R ../worker/views/themes/worker/assets ./public/admin/

release: clean template assets
	@cp -R ./public ./dist/
	@#go-bindata -nomemcopy ../qor/admin/views/...
	@echo "building release ${BIN_NAME} ${VERSION}"
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w -X main.BuildTime=${CUR_TIME} -X main.Version=${VERSION} -X main.GitHash=${GIT_COMMIT}' -o ./dist/$(BIN_NAME) main.go
	@echo "building release ${BIN_NAME_CLI} ${VERSION}"
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w -X main.BuildTime=${CUR_TIME} -X main.Version=${VERSION} -X main.GitHash=${GIT_COMMIT}' -o ./dist/$(BIN_NAME_CLI) cli.go
	@chmod 0755 ./dist/$(BIN_NAME_CLI)

clean:
	@test ! -e ./${BIN_NAME} || rm ./${BIN_NAME}
	@git gc --prune=0 --aggressive
	@find . -name "*.orig" -type f -delete
	@find . -name "*.log" -type f -delete
	@test ! -e ./dist || rm -R ./dist

seed:
	@echo "...............................................................\n"
	@echo $(PROJECT_NAME) seed
	@echo ...............................................................
	@go run db/seeds/main.go

run:
	@echo "...............................................................\n"
	@echo Project: $(PROJECT_NAME)
	@echo Open in browser:
	@echo	"	 http://localhost:7000/\n"
	@echo ...............................................................
	@QORCONFIG=config/database.dev.yml go run main.go

test:
	@go test -v ./...
	@#API_PATH=$(PROJECT_DIR) ginkgo -v -r

build: clean
	@echo "Building ${BIN_NAME} ${VERSION}"
	@CGO_ENABLED=0 go build -a -tags netgo -ldflags '-w -X main.BuildTime=${CUR_TIME} -X main.Version=${VERSION} -X main.GitHash=${GIT_COMMIT}' -o $(BIN_NAME) main.go
	@echo "Building ${BIN_NAME_CLI} ${VERSION}"
	@CGO_ENABLED=0 go build -a -tags netgo -ldflags '-w -X main.BuildTime=${CUR_TIME} -X main.Version=${VERSION} -X main.GitHash=${GIT_COMMIT}' -o $(BIN_NAME_CLI) cli.go


cli: clean
	@echo "Building cli ${VERSION}"
	@CGO_ENABLED=0 go build -a -tags netgo -ldflags '-w -X main.BuildTime=${CUR_TIME} -X main.Version=${VERSION} -X main.GitHash=${GIT_COMMIT}' -o $(BIN_NAME_CLI) cli.go

docs:
	godoc -http=:6060 -index

