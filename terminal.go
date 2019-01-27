package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var quotes = []rune{'"', '\''}
var quotes_unexpandable = []rune{'\''}

type Context struct {
	current_dir string
	output      *os.File
	input       *os.File
	error       *os.File
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

	for true {
		fmt.Fprint(ctx.output, get_head_line(&ctx))
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(ctx.error, err)
		}
		err = parse_command(&ctx, input)
		if err != nil {
			fmt.Fprintln(ctx.error, err)
		}
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

func parse_pipe(ctx *Context, commands string) error {
	return nil
}

func parse_command(ctx *Context, command string) error {

	input := split_and_expand_input(ctx, []rune(strings.TrimSuffix(command, "\n")), quotes, quotes_unexpandable)

	//todo: chain the | here

	//internal command check
	switch input[0] {
	case "exit":
		return internal_exit(ctx)
	case "cd":
		if len(input) > 1 {
			return internal_cd(ctx, input[1])
		} else {
			return internal_cd(ctx, "")
		}
	case "pwd":
		return internal_pwd(ctx)
	}
	return exec_command(ctx, input)
}

func split_and_expand_input(c *Context, command []rune, delimiters []rune, unepandables []rune) []string {
	block := false
	expandable := false
	var delimiter rune
	current := ""
	res := make([]string, 0)
	for index := 0; index < len(command); index++ {
		letter := command[index]
		if contains(delimiters, letter) && (index == 0 || command[index-1] == ' ') {
			current = ""
			expandable = !contains(unepandables, letter)
			//start block
			block = true
			delimiter = letter
			continue
		}
		if letter == delimiter && (index == len(command)-1 || command[index+1] == ' ') {
			//end block
			res = append(res, current)
			current = ""
			continue
		}

		//cut by word
		if !block && letter == ' ' {
			res = append(res, current)
			current = ""
			continue
		}

		//expand the ~
		if !block && letter == '~' && (index == 0 || command[index-1] == ' ') {
			command = []rune(strings.Replace(string(command), "~", get_user_dir(c), 1))
			index--
			continue
		}

		//todo : integrate the '|' treatement here
		if (!block || expandable) && current == "" && letter == '*' {
			//todo: expand * and */** here
			if index < len(command)-3 && string(command[index:index+4]) == "*/**" {
				//*/** here
				res = append(res, get_content_folder_recursive(c, c.current_dir)...)
				index += 3
				continue
			} else {
				//* here
				res = append(res, get_content_folder(c.current_dir)...)
				continue
			}
		}
		current += string(letter)
	}
	if current != "" {
		res = append(res, current)
	}
	res = delete_empty_elements(res)
	return res
}

//todo manage error (404 ...)
func exec_command(ctx *Context, input []string) error {
	fork := []rune(input[len(input)-1])[len([]rune(input[len(input)-1]))-1] == '&'
	if fork { //remove the & from the last arg .... this isn't pretty
		if input[len(input)-1] == "&" {
			input = input[:len(input)-1]
		} else {
			input = append(input[:len(input)-1], string([]rune(input[len(input)-1])[:len(input[len(input)-1])-1]))
		}
	}
	var cmd *exec.Cmd
	if len(input) == 1 {
		cmd = exec.Command(input[0], make([]string, 0)...)
	} else {
		cmd = exec.Command(input[0], input[1:]...)
	}
	cmd.Stderr = ctx.error
	cmd.Stdout = ctx.output
	cmd.Stdin = ctx.input
	if !fork {
		return cmd.Run()
	} else {
		err := cmd.Start()
		fmt.Fprintln(ctx.output, "New process started with the pid ", cmd.Process.Pid)
		return err
	}
}
