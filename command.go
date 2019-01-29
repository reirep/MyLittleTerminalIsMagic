package main

import (
	"os"
	"os/exec"
)

type command struct {
	sync    bool
	exec    string
	args    []string
	context *Context
	run     *exec.Cmd
}

func new_command(sync bool, exe string, args []string, ctx *Context) command {
	command := command{
		sync:    sync,
		exec:    exe,
		args:    args,
		context: ctx,
	}
	if len(command.args) == 0 {
		command.run = exec.Command(exe, make([]string, 0)...)
	} else {
		command.run = exec.Command(exe, args[:]...)
	}
	command.run.Stdin = ctx.input
	command.run.Stdout = ctx.output
	command.run.Stderr = ctx.error
	return command
}

func (c command) start() error {
	if c.sync {
		return c.run.Run()
	} else {
		return c.run.Start()
	}
}

func (c command) wait_end() error {
	return c.run.Wait()
}

func (c command) send_signint() error {
	process, err := os.FindProcess(c.run.Process.Pid)
	if err != nil {
		return err
	}
	process.Signal(os.Interrupt)
	return nil
}
