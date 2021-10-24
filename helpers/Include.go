package helpers

import (
	"fmt"
	"path/filepath"
)

func Include(path string) []string {
	files, err := filepath.Glob("admin/views/templates/*.html")
	if err != nil {
		fmt.Print(err)

	}
	path_Files, err := filepath.Glob("admin/views/" + path + "/*.html")
	if err != nil {
		fmt.Print(err)
	}

	for _, file := range path_Files {
		files = append(files, file)
	}
	return files
}
