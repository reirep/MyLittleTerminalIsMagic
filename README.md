# My little command line is magic
Just another command line interpreter, I may work from time to time on it.

## Currently implemented
- go to a dir: `cd <dir>`
- go to the home dir: `cd`
- show current folder: `pwd`
- display `#` when it is a root terminal; `$` in any other case
- fork with `&` at the end of the command
- `"` block
- `'` block
- expand the ~
- list of the files of the current dir with `*` (except when the `*` is in a block)
- list all the files and the subfiles with `*/**` (except when `*/**` is in a block)

## To fix / Currently happening

## TODO
- some way to manage the history
- auto expand all the know command (use the history ?) use tab
- piping with `|` (use a pipe to do that https://golang.org/pkg/io/#Pipe)
- add support for the left arrow key and right arrow key
- add support for the end and home key and cursor wandeling arround
