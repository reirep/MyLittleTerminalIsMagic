package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var rune_pipe = '|'
var pipe_stderr = false
var quotes = []rune{'"', '\''}
var quotes_unexpandable = []rune{'\''}

type Context struct {
	current_dir string
	output      io.Writer
	input       io.Reader
	error       io.Writer
}

func (c Context) clone() Context {
	ctx := Context{
		current_dir: c.current_dir,
		output:      c.output,
		input:       c.input,
		error:       c.error,
	}
	return ctx
}

func main() {
	ctx := Context{
		output: os.Stdout,
		input:  os.Stdin,
		error:  os.Stderr,
	}
	ctx.current_dir = get_current_dir(&ctx)

	reader := bufio.NewReader(ctx.input)

	for {
		fmt.Fprint(ctx.output, get_head_line(&ctx))
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(ctx.error, err)
		}
		parse_pipe(&ctx, input)
	}
}

func get_head_line(c *Context) (res string) {
	user := get_username(c)
	res += user
	res += "@"
	res += get_hostname(c)
	res += " ("
	res += c.current_dir
	res += ")\n"
	if user == "root" {
		res += "#"
	} else {
		res += "$"
	}
	res += " "
	return res
}

func parse_pipe(ctx *Context, commands string) {
	input := split_and_expand_input(ctx, []rune(strings.TrimSuffix(commands, "\n")))

	//chain the commands
	contexts := make([]*Context, len(input))
	for i := 0; i < len(contexts); i++ {
		context := ctx.clone()
		contexts[i] = &context
	}
	for i := 0; i < len(contexts)-1; i++ {
		reader, writer := io.Pipe()
		contexts[i].output = writer
		contexts[i+1].input = reader
	}

	for i := 0; i < len(input); i++ {
		if i != len(input)-1 { //wait for the last part of the pipe
			go parse_command(contexts[i], input[i])
		} else {
			parse_command(contexts[i], input[i])
		}
	}
}

func parse_command(ctx *Context, command []string) {
	//internal command check
	switch command[0] {
	case "exit":
		internal_exit(ctx)
		return
	case "cd":
		if len(command) > 1 {
			internal_cd(ctx, command[0])
			return
		} else {
			internal_cd(ctx, "")
			return
		}
	case "pwd":
		internal_pwd(ctx)
		return
	}
	exec_command(ctx, command)
}

func exec_command(ctx *Context, input []string) command {
	sync := is_command_sync(input)
	if !sync { //remove the & if needed
		input = remove_last_letter_command(input)
	}
	input = delete_empty_elements_array_string(input)

	command := new_command(sync, input[0], input[1:], ctx)
	err := command.start()
	if err != nil {
		fmt.Fprintln(ctx.error, err)
	}
	return command
}

func split_and_expand_input(c *Context, command []rune) [][]string {
	index_current := 0
	block := false
	expandable := false
	var delimiter rune
	current := ""
	res := make([][]string, 1)
	res[0] = make([]string, 0)

	for index := 0; index < len(command); index++ {
		letter := command[index]

		//manage blocks
		if contains(quotes, letter) && (index == 0 || command[index-1] == ' ') {
			current = ""
			expandable = !contains(quotes_unexpandable, letter)
			//start block
			block = true
			delimiter = letter
			continue
		}
		if letter == delimiter && (index == len(command)-1 || command[index+1] == ' ') {
			//end block
			res[index_current] = append(res[index_current], current)
			current = ""
			continue
		}

		//cut by word
		if !block && letter == ' ' {
			res[index_current] = append(res[index_current], current)
			current = ""
			continue
		}

		//expand the ~
		if !block && letter == '~' && (index == 0 || command[index-1] == ' ') {
			command = []rune(strings.Replace(string(command), "~", get_user_dir(c), 1))
			index-- //could do += len(get_user_Dir(c)) -1
			continue
		}

		//manage the pipe cutting
		if !block && letter == rune_pipe {
			if current != "" {
				res[index_current] = append(res[index_current], current)
			}
			res = append(res, make([]string, 0))
			index_current++
			continue
		}

		if (!block || expandable) && current == "" && letter == '*' {
			//expand * and */** here
			if index < len(command)-3 && string(command[index:index+4]) == "*/**" {
				//*/** here
				res[index_current] = append(res[index_current], get_content_folder_recursive(c, c.current_dir)...)
				index += 3
				continue
			} else {
				//* here
				res[index_current] = append(res[index_current], get_content_folder(c.current_dir)...)
				continue
			}
		}
		current += string(letter)
	}
	if current != "" {
		res[index_current] = append(res[index_current], current)
	}
	res = delete_empty_elements_array_array(res)
	return res
}
