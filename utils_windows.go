// +build windows

package main

import "regexp"

func is_root_path(path string) bool {
	if len(path) < 3 {
		return false
	}
	re := regexp.MustCompile(`(?m)^[a-zA-z]:\\.*$`)
	return re.MatchString(path)
}
