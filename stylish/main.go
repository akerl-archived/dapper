package stylish

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

const (
	// DefaultFile defines the standard location for the packed JSON file
	DefaultFile = "rendered/stylish.json"
	// DefaultDir defines the standard location for the unpacked directory
	DefaultDir = "src"

	configFile = "config.json"
	cssFile    = "style.css"
	sectionDir = "sections"
)

var cleanNameRegexp = regexp.MustCompile(`[^\w]+`)

func readJSONFile(object interface{}, pathSegments ...string) error {
	data, err := readTextFile(pathSegments...)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &object)
}

func readTextFile(pathSegments ...string) ([]byte, error) {
	path := filepath.Join(pathSegments...)
	return ioutil.ReadFile(path) // #nosec
}

func writeJSONFile(data interface{}, pathSegments ...string) error {
	buffer, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return writeTextFile(buffer, pathSegments...)
}

func writeTextFile(data []byte, pathSegments ...string) error {
	path := filepath.Join(pathSegments...)
	return ioutil.WriteFile(path, data, 0600)
}

func cleanDir(path string) error {
	if err := os.RemoveAll(path); err != nil {
		return err
	}
	return os.MkdirAll(path, 0700)
}
