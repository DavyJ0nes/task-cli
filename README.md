# Task CLI

[![Go Report Card](https://goreportcard.com/badge/github.com/DavyJ0nes/task-cli)](https://goreportcard.com/report/github.com/DavyJ0nes/task-cli)

## Description

Task is a simple CLI tool to help manage your life.

It uses the [cobra package](https://github.com/spf13/cobra) to help speed up development.

It was created as part of a tutorial on [Gophercises](https://gophercises.com/exercises/task). You can check out the original repository [here](https://github.com/gophercises/task)

Thanks to [Jon Calhoun](https://twitter.com/joncalhoun) for creating this great learning resource.

## Installation

### go install

If you have your Go workspace set up correctly then you are able to install task-cli to your `$GOPATH/bin` directory by using `go install`. You may need to run `go get` before this if there are dependency issues, though `go install` should take care of most of this for you.

### Docker

If you would like you can run task-cli within a docker contianer. As part of the Makefile you can run `make install` this will perform the following actions

- Create a statically linked binary
- Create a docker image with the go binary within
- Create a docker volume to persist the bolt database file
- Create a wrapper bash script within `$HOME/bin` called `taskd` that will run the container

You can also forego using the wrapper bash script and run the container yourself with the following commands:

```shell
docker volume create task
docker run -it --rm -v task:/app/.tasks davyj0nes/task:latest
```

## Usage

Basic Usage instructions.

```none
Task is a CLI todo tool

Usage:
  task [command]

Available Commands:
  add         adds a command to the task list
  do          completes a task
  help        Help about any command
  list        Lists the current open tasks

Flags:
  -h, --help   help for task

Use "task [command] --help" for more information about a command.
```

## TODO

- [x] Add Add Command
- [x] Add List Command
- [x] Add Do Command
- [x] Add database (boltdb)
- [x] Complete tutorial code
- [x] Ensure tool can be run from within Docker
- [ ] Add metadata to Tasks
- [ ] Add completed flag to List command
- [ ] Fork logic to create API backend for web service

## License

[MIT](./LICENSE)
