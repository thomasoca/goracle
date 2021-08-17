# goracle
A CLI app to watch a directory and doing something, powered by golang and fsnotify
## Usage
Work is still on progress to make it a proper CLI (see [To-Do](#to-do) section for more example), but overall the usage will be:

```shell
goracle [options] COMMAND ARGS
```

There are some options available:
```shell
Options:
  -dir string
        Directory to watch, set to current working directory as default (default "/home/thomasoca/projects/goracle")
  -e string
        Event to notify, select between create, write, remove, rename, chmod (default "create")
  -nb
        Set execution mode to non-blocking
  -pattern string
        File pattern to notify (default "*")
```

NOTE: Separate multiple ARGS by space

## Usage example 
Run a python program over a write event of .csv file

```shell
goracle -e write -dir /home/users/folder -nb -pattern *.csv python example.py input_argument
```

The event file name will always be passed to the last ARGS by default

## To-do
- [ ] Write more tests to cover most of the statements
- [ ] Create the makefile for CLI installation
- [ ] Re-thinking about the non-blocking option
