all: run

# Example Makefile for building and deploying Go applications with Docker

# To build and run the docker container locally, run:
# $ make

# or to publish the :latest version to the specified registry as :1.0.0, run:
# $ make publish version=1.0.0

#### VARIABLES ####
username = davyj0nes
app_name = task

binary_version = 0.0.1
image_version ?= latest

go_version ?= 1.10

git_hash = $(shell git rev-parse HEAD | cut -c 1-6)
build_date = $(shell date -u '+%Y-%m-%d_%I:%M:%S%p')

.PHONY: run binary image run_image test clean

#### COMMANDS ####
run:
	$(call blue, "# Running App...")
	@docker run -it --rm -v "$(GOPATH)":/go -v "$(CURDIR)":/go/src/app -w /go/src/app golang:${go_version} go run main.go

binary:
	$(call blue, "# Building Golang Binary...")
	@docker run --rm -v "$(CURDIR)":/go/src/app -w /go/src/app golang:${go_version} sh -c 'go get && CGO_ENABLED=0 GOOS=linux go build -a -tags netgo --installsuffix netgo -o ${app_name}'

image: binary
	$(call blue, "# Building Docker Image...")
	@docker build --label APP_VERSION=${binary_version} --label BUILT_ON=${build_date} --label GIT_HASH=${git_hash} -t ${username}/${app_name}:${image_version} .
	@$(MAKE) clean

run_image: image
	$(call blue, "# Running Docker Image Locally...")
	@docker run -it --rm --name ${app_name} ${username}/${app_name}:${image_version} 

test:
	$(call blue, "# Testing Golang Code...")
	@docker run --rm -it -v "$(GOPATH):/go" -v "$(CURDIR)":/go/src/app -w /go/src/app golang:${go_version} sh -c 'go test -v' 

clean: 
	@rm -f ${app_name} 

#### FUNCTIONS ####
define blue
	@tput setaf 4
	@echo $1
	@tput sgr0
endef
