// +build !windows

package main

func is_root_path(path string) bool {
	if len(path) < 1 {
		return false
	}
	return []rune(path)[0] == '/'
}
