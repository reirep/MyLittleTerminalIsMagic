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

func main() {
	ctx := Context{
		output: os.Stdout,
		input:  os.Stdin,
		error:  os.Stderr,
	}
	ctx.current_dir = get_current_dir(&ctx)

	reader := bufio.NewReader(ctx.input)

	for {
		fmt.Fprint(ctx.output, ctx.current_dir, "> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(ctx.error, err)
		}
		parse_command(&ctx, input)
	}
	fmt.Println("Hello world !")
}

func parse_command(ctx *Context, command string) error {
	input := strings.Split(strings.TrimSuffix(command, "\n"), " ")

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

	cmd := exec.Command(input[0], input[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}
