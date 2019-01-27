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
- list of the files of the current directory with `*` (except when the `*` is in a block)
- list all the files and the sub-files with `*/**` (except when `*/**` is in a block)

## To fix / Currently happening

## TODO
- use some stringbuilder when needed
- manage ctrl + c to cut the underlying command
- check that all the code handle utf8 correctly #rune
- some way to manage the history
- auto expand all the know command (use the history ?) use tab
- piping with `|` (use a pipe to do that https://golang.org/pkg/io/#Pipe)
- add support for the left arrow key and right arrow key
- add support for the end and home key and cursor wandering around
