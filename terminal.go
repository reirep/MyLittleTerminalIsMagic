package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

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

	for {
		fmt.Fprint(ctx.output, get_head_line(&ctx))
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(ctx.error, err)
		}
		parse_command(&ctx, input)
	}
}

func get_head_line(c *Context) (res string) {
	user := get_username(c)
	res += user
	res += "@"
	res += get_hostname(c)
	res += " ("
	res += get_current_dir(c)
	res += ")\n"
	if user == "root" {
		res += "#"
	} else {
		res += "$"
	}
	res += " "
	return res
}

func parse_command(ctx *Context, command string) error {
	input := strings.Split(strings.TrimSuffix(command, "\n"), " ")

	//internal command test
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

func exec_command(ctx *Context, input []string) error {
	fork := input[len(input)-1] == "&"
	if fork {
		input = input[:len(input)-1]
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
