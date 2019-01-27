package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

func internal_exit(c *Context) error {
	os.Exit(0)
	return nil
}

//todo save complete path in the current dir var
func internal_cd(c *Context, emplacement string) error {
	var new_dir string
	if emplacement != "" {
		new_dir = emplacement
	} else {
		new_dir = get_user_dir(c)
	}
	if !is_root_path(new_dir) {
		new_dir = path.Join(c.current_dir, new_dir)
	}

	err := os.Chdir(new_dir)
	if err != nil {
		return err
	} else {
		c.current_dir = filepath.Clean(new_dir)
		return nil
	}
}

func internal_pwd(c *Context) error {
	fmt.Fprintln(c.output, c.current_dir)
	return nil
}
