package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

func internal_exit(c *Context) {
	os.Exit(0)
}

func internal_cd(c *Context, emplacement string) {
	var new_dir string
	if emplacement == "" {
		new_dir = get_user_dir(c)
	} else {
		new_dir = emplacement
	}

	if !is_root_path(new_dir) {
		new_dir = path.Join(c.current_dir, new_dir)
	}

	err := os.Chdir(new_dir)
	if err != nil {
		fmt.Fprintln(c.error, err)
	} else {
		c.current_dir = filepath.Clean(new_dir)
	}
}

func internal_pwd(c *Context) {
	fmt.Fprintln(c.output, c.current_dir)
}
