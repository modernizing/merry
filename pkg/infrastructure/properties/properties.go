package properties

import (
	"archive/zip"
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type AppConfigProperties map[string]string

type myCloser interface {
	Close() error
}
// check is a helper function which streamlines error checking
func check(e error) {
	if e != nil {
		panic(e)
	}
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


func ReadPropertiesFromZip(zipFile string) (AppConfigProperties, error) {
	config := AppConfigProperties{}

	zf, err := zip.OpenReader(zipFile)
	check(err)
	defer closeFile(zf)

	var hasProperties = false
	for _, file := range zf.File {
		if strings.HasSuffix(file.Name, "pom.properties") {
			hasProperties = true
			all := readAll(file)
			content := string(all)
			lines := strings.Split(content, "\n")
			for _, line := range lines {
				buildConfigByLine(line, config)
			}
		}
	}

	if !hasProperties {
		return nil, err
	}

	return config, err
}

// https://stackoverflow.com/questions/40022861/parsing-values-from-property-file-in-golang
func ReadPropertiesFile(filename string) (AppConfigProperties, error) {
	config := AppConfigProperties{}

	if len(filename) == 0 {
		return config, nil
	}
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		buildConfigByLine(line, config)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return config, nil
}

func buildConfigByLine(line string, config AppConfigProperties) {
	if equal := strings.Index(line, "="); equal >= 0 {
		if key := strings.TrimSpace(line[:equal]); len(key) > 0 {
			value := ""
			if len(line) > equal {
				value = strings.TrimSpace(line[equal+1:])
			}
			config[key] = value
		}
	}
}
