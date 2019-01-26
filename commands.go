package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

func internal_exit(c *Context) error {
	os.Exit(0)
	return nil
}

//internal commands

func internal_cd(c *Context, emplacement string) error {
	var new_dir string
	if emplacement != "" {
		new_dir = emplacement
	} else {
		new_dir = get_user_dir(c)
	}
	c.current_dir = new_dir
	return os.Chdir(new_dir)
}

func internal_pwd(c *Context) error {
	fmt.Fprintln(c.output, c.current_dir)
	return nil
}

// utils

func get_user_dir(c *Context) string {
	usr, err := user.Current()
	if err != nil {
		fmt.Fprintln(c.error, err)
		return ""
	}
	return usr.HomeDir
}

func get_current_dir(c *Context) string {
	ex, err := os.Executable()
	if err != nil {
		fmt.Fprintln(c.error, err)
		return ""
	}
	return filepath.Dir(ex)
}
