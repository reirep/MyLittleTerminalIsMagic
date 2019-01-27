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

func get_content_folder(foldert string) []string {
	return []string{"lorem ipsum"} //todo
}

func get_content_folder_recursive(c *Context, folder string) []string {
	var files []string
	err := filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		fmt.Fprintln(c.error, err)
		return nil
	}
	return files
}

func contains(letters []rune, letter rune) bool {
	for _, current := range letters {
		if current == letter {
			return true
		}
	}
	return false
}

func delete_empty_elements(in [][]string) [][]string {
	for i := 0; i < len(in); i++ {
		if len(in[i]) == 0 {
			in = append(in[:i], in[i+1:]...)
			i--
			continue
		}
		in[i] = delete_empty_elements_simple(in[i])
	}
	return in
}

func delete_empty_elements_simple(in []string) []string {
	for i := 0; i < len(in); i++ {
		if in[i] == "" {
			in = append(in[:i], in[i+1:]...)
			i--
		}
	}
	return in
}

//string utils that account for the utf8 encoding
func charAt(str string, index int) rune {
	return []rune(str)[index]
}

func subStr(str string, start int, end int) string {
	return string(([]rune(str))[start:end])
}
