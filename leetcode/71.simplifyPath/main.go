package main

import (
	"fmt"
	"strings"
)

func simplifyPath(path string) string {
	dirs := strings.FieldsFunc(path, func(c rune) bool {
		return c == '/'
	})

	simplifiedDirs := make([]string, 0, len(dirs))
	for index := 0; index < len(dirs); index++ {
		switch dirs[index] {
		case ".":
			continue
		case "..":
			if len(simplifiedDirs) > 0 {
				simplifiedDirs = simplifiedDirs[:len(simplifiedDirs)-1]
			}
		default:
			simplifiedDirs = append(simplifiedDirs, dirs[index])
		}
	}
	return "/" + strings.Join(simplifiedDirs, "/")
}

func main() {
	//path := "/home/user/Documents/../Pictures"
	path := "/.../a/../b/c/../d/./"
	fmt.Println(simplifyPath(path))
}
