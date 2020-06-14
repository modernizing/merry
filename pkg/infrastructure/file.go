package infrastructure

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
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

func GetJarFilesByPath(path string, excludeSource bool) []string {
	var jarFileFilter = func(path string) bool {
		return strings.HasSuffix(path, ".jar")
	}
	if excludeSource {
		jarFileFilter = func(path string) bool {
			return strings.HasSuffix(path, ".jar") && !strings.Contains(path, ".source")
		}
	}

	jarPaths := GetFilesByFilter(path, jarFileFilter)
	return jarPaths
}

func GetJarPath(path string) string {
	separator := string(os.PathSeparator)
	split := strings.Split(path, separator)
	purePath := split[:len(split)-1]
	result := strings.Join(purePath, separator)
	return result + separator
}

func Copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}
