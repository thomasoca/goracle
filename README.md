# goracle
A CLI app to watch a directory and doing something, powered by golang and fsnotify
## Usage [WIP]
Work is still on progress to make it a proper CLI, but overall the usage will be:

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
goracle -e write -dir /home/users/projects -pattern *.csv python test.py
```

The event file name will always be passed to the last ARGS by default
