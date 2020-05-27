export GO111MODULE=on
# Get the current full sha from git
GITSHA:=$(shell git rev-parse --short HEAD)
# Get the current local branch name from git (if we can, this may be blank)
GITBRANCH:=$(shell git symbolic-ref --short HEAD 2>/dev/null)
GOFMT_FILES?=$$(find . -not -path "./vendor/*" -name "*.go")
COVER_DIR=./.cover
COVER_FILE="${COVER_DIR}/coverall.out"
COVER_HTML="${COVER_DIR}/coverall.html"
PACKAGES=$(shell go list ./... | grep -v '/vendor/')
VERSION?=$(shell cat ./VERSION)-${GITSHA}
DOCKER_IMAGE?=deploy-service:${VERSION}
DOCKER_IMAGE_LATEST=deploy-service:latest
GOOS?=darwin
GOMODFLAG?=-mod=vendor # set modules mode during development. Unset this variable when building in CI to reduce network overhead.
CGO_ENABLED?=0
GOARCH?=amd64
GOSUMDB=off



.PHONY: build install doc fmt lint run test vet setup docker-image docker-image clean all linux-bin
default: all

all: setup deps generate lint vet test install

setup: $(GOPATH)/bin/golint $(GOPATH)/bin/goimports $(GOPATH)/bin/swagger
	
$(GOPATH)/bin/golint:
	go get golang.org/x/lint/golint


$(GOPATH)/bin/goimports:
	go get golang.org/x/tools/cmd/goimports

$(GOPATH)/bin/swagger:
	go get github.com/go-swagger/go-swagger/cmd/swagger

deps:
	go mod download
	go mod vendor -v

test:
	@echo "tests..."
	@go test ${GOMODFLAG} -v -timeout=2m ${PACKAGES}

cover-html:
	rm -rf ${COVER_DIR}
	mkdir ${COVER_DIR}
	echo "mode: count" > ${COVER_FILE}
	$(foreach pkg, ${PACKAGES},\
		go test -coverprofile=${COVER_DIR}/cover-pkg.pkg.out -covermode=count ${pkg};\
		grep -h -v "^mode:" ${COVER_DIR}/*.pkg.out >> ${COVER_FILE} || true;\
	)
	go tool cover -html=${COVER_FILE} -o ${COVER_HTML}

format: 
	@echo "formatting..."
	@goimports ${GOFMT_FILES} 1>/dev/null

generate:
	@echo "generating..."
	@go generate

build: format generate
	@echo "building..."
	@go build ${GOMODFLAG} ${PACKAGES}

install:
	@echo "building and installing..."
	@GOOS=${GOOS} CGO_ENABLED=${CGO_ENABLED} GOARCH=${GOARCH} go install --installsuffix cgo --ldflags="-w  -s" ${GOMODFLAG} ./cmd/$*

docker-image:
	echo "Building docker image: ${DOCKER_IMAGE}"
	docker build --build-arg GITSHA=${GITSHA} -t ${DOCKER_IMAGE} -f Dockerfile ${PWD}
	echo "Tagging image to: ${DOCKER_IMAGE_LATEST}"
	docker tag ${DOCKER_IMAGE} ${DOCKER_IMAGE_LATEST}

doc:
	@godoc -http=:6060 -index
	

# https://github.com/golang/lint
# go get github.com/golang/lint/golint
lint:
	@echo "linting..."
	@golint ${PACKAGES}

run: build
	$(GOPATH)/bin/deploy-service -c $(config)

vet:
	@echo "govet..."
	@go vet ${GOMODFLAG} -v ${PACKAGES}

clean:
	go clean -i -x -r
