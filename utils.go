package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

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

func get_hostname(c *Context) string {
	host, err := os.Hostname()
	if err != nil {
		fmt.Fprintln(c.error, err)
		return ""
	}
	return host
}

func get_username(c *Context) string {
	usr, err := user.Current()
	if err != nil {
		fmt.Fprintln(c.error, err)
		return ""
	}
	return usr.Username
}

func get_env(c *Context) []string {
	return os.Environ()
}

func get_env_var(c *Context, varEnv string) string {
	return os.Getenv(varEnv)
}

func contains(letters []rune, letter rune) bool {
	for _, current := range letters {
		if current == letter {
			return true
		}
	}
	return false
}

//string utils that account for the utf8 encoding
func charAt(str string, index int) rune {
	return []rune(str)[index]
}

func subStr(str string, start int, end int) string {
	return string(([]rune(str))[start:end])
}
