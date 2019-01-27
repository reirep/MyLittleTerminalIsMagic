# My little command line is magic
This is just another command line interpreter, I'm working from time to time on it :) . At the moment, don't expect much, this is what EA would call a finished game (= an _alpha_ ).
To see a list of what it's currently capable of doings, look at the list bellow

⚠ - This terminal is currently working with the simple command but the pipes are **broken**: they will execute and pass data around but won't end.

⚠ - This terminal cannot handle yet the sigint (so, no way to interrupt a command).

## License
This project is under the MIT license - license file coming soon.

All contribution are welcome. By making a contribution, you're giving me the right to the contributed code (to allow this project to be easier to maintain and not having to ask everyone their comment, opinion and feelings on a single change)

## Roadmap
### Currently implemented
- execute simple commands
- go to a dir: `cd <dir>`
- go to the home dir: `cd`
- show current folder: `pwd`
- display `#` when it is a root terminal; `$` in any other case
- fork the command if is `&` at the end of the command
- `"` block
- `'` block
- expand the ~
- list of the files of the current directory with `*` (except when the `*` is in a block)
- list all the files and the sub-files with `*/**` (except when `*/**` is in a block)

### To fix / Currently happening
- piping with `|` -> the pipe doesn't stop gracefully, to fix

### TODO
- redirect the stream with the > notation
- redirect the stream themselves with &1 and &2
- use some stringbuilder when needed
- manage ctrl + c to cut the underlying command -> https://gobyexample.com/signals
- check that all the code handle utf8 correctly #rune
- some way to manage the history
- auto expand all the know command (use the history ?) use tab
- add support for the left arrow key and right arrow key
- add support for the end and home key and cursor wandering around

### TODO non code-related
- Create a license file
