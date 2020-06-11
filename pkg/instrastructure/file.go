package instrastructure

import (
	"os"
	"path/filepath"
)

func GetFilesByFilter(root string, filter func(path string) bool) []string {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if filter(path) {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return files
}
