package controller

import "strings"

func isExcludeFolder(path string, path_ignore []string) bool {

	for _, exclude := range path_ignore {
		if strings.Contains(path, exclude) {
			return true
		}
	}

	return false
}
