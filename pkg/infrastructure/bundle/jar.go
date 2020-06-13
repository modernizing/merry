package bundle

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type myCloser interface {
	Close() error
}

// closeFile is a helper function which streamlines closing
// with error checking on different file types.
func closeFile(f myCloser) {
	err := f.Close()
	check(err)
}

// readAll is a wrapper function for ioutil.ReadAll. It accepts a zip.File as
// its parameter, opens it, reads its content and returns it as a byte slice.
func readAll(file *zip.File) []byte {
	fc, err := file.Open()
	check(err)
	defer closeFile(fc)

	content, err := ioutil.ReadAll(fc)
	check(err)

	return content
}

// check is a helper function which streamlines error checking
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ListJarFiles(path string) []zip.File {
	zf, err := zip.OpenReader(path)
	check(err)
	defer closeFile(zf)

	var files []zip.File
	for _, file := range zf.File {
		files = append(files, *file)
	}
	return files
}

func GetFileFromJar(zipFile string, fileName string) (string, string, error) {
	zf, err := zip.OpenReader(zipFile)
	check(err)
	defer closeFile(zf)

	for _, file := range zf.File {
		if strings.HasSuffix(file.Name, fileName) {
			all := readAll(file)
			content := string(all)
			return file.Name, content, nil
		}
	}

	return "", "", err
}


// https://stackoverflow.com/questions/20357223/easy-way-to-unzip-file-with-golang
func UnZip(src, dest string) ([]string, error) {
	var fileNames []string
	r, err := zip.OpenReader(src)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := r.Close(); err != nil {
			panic(err)
		}
	}()

	os.MkdirAll(dest, 0755)

	// Closure to address file descriptors issue with all the deferred .Close() methods
	extractAndWriteFile := func(f *zip.File) error {
		rc, err := f.Open()
		if err != nil {
			fmt.Println(err)
			return err
		}
		defer func() {
			if err := rc.Close(); err != nil {
				panic(err)
			}
		}()

		path := filepath.Join(dest, f.Name)

		if f.FileInfo().IsDir() {
			os.MkdirAll(path, 0755)
		} else {
			os.MkdirAll(filepath.Dir(path), 0755)
			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				fmt.Println(err)
				return err
			}
			defer func() {
				if err := f.Close(); err != nil {
					panic(err)
				}
			}()

			_, err = io.Copy(f, rc)
			if err != nil {
				fmt.Println(err)
				return err
			}
			fileNames = append(fileNames, f.Name())
		}
		return nil
	}

	for _, f := range r.File {
		err := extractAndWriteFile(f)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}

	return fileNames, nil
}
