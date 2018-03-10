# Task CLI

## Description

Task is a simple CLI tool to help manage your life.

It uses the [cobra package](https://github.com/spf13/cobra) to help speed up development.

It was created as part of a tutorial on [Gophercises](https://gophercises.com/exercises/task). You can check out the original repository [here](https://github.com/gophercises/task)

Thanks to [Jon Calhoun](https://twitter.com/joncalhoun) for creating this great learning resource.

## Usage

Basic Usage instructions.

```shell
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
- [ ] Add database (boltdb)
- [ ] Complete tutorial code
- [ ] Allow easy swapping of database backends for tool with interfaces
- [ ] Update Makefile to create releases for relevant OSes.
- [ ] Ensure tool can be run from within Docker

## License

MIT
