package stylish

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
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

var cleanNameRegexp = regexp.MustCompile("[^\\w]+")

type section struct {
	URLs        []string `json:"urls"`
	URLPrefixes []string `json:"urlPrefixes"`
	Domains     []string `json:"domains"`
	Regexps     []string `json:"regexps"`
	Code        string   `json:"code"`
}

// Style describes a Stylish theme
type Style struct {
	Sections    []section `json:"sections"`
	URL         string    `json:"url"`
	UpdateURL   string    `json:"updateUrl"`
	MD5URL      string    `json:"md5Url"`
	OriginalMD5 string    `json:"originalMd5"`
	Name        string    `json:"name"`
	Enabled     bool      `json:"enabled"`
	ID          int       `json:"id"`
}

// StyleSet is a slice of Styles
type StyleSet []Style

// ReadFromFile reads a StyleSet from a packed file
func ReadFromFile(file string) (StyleSet, error) {
	ss := StyleSet{}

	data, err := ioutil.ReadFile(file)
	if err != nil {
		return ss, err
	}

	err = json.Unmarshal(data, &ss)
	return ss, err
}

// WriteToFile writes a StyleSet to a packed file
func (ss StyleSet) WriteToFile(file string) error {
	buffer, err := json.MarshalIndent(ss, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(file, buffer, 0600)
}

// ReadFromDir reads a StyleSet from an unpacked dir
func ReadFromDir(dir string) (StyleSet, error) {
	ss := StyleSet{}

	sectionGlob := filepath.Join(dir, "*", sectionDir)
	sectionDirs, err := filepath.Glob(sectionGlob)
	if err != nil {
		return ss, err
	}
	for _, x := range sectionDirs {
		styleDir := filepath.Dir(x)
		s, err := readStyleFromDir(styleDir)
		if err != nil {
			return ss, err
		}
		ss = append(ss, s)
	}
	return ss, nil
}

func readStyleFromDir(dir string) (Style, error) {
	s := Style{}

	configPath := filepath.Join(dir, configFile)
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return s, err
	}
	err = json.Unmarshal(data, &s)
	if err != nil {
		return s, err
	}

	sectionGlob := filepath.Join(dir, sectionDir, "*")
	sectionDirs, err := filepath.Glob(sectionGlob)
	if err != nil {
		return s, err
	}
	for _, x := range sectionDirs {
		sect, err := readSectionFromDir(x)
		if err != nil {
			return s, err
		}
		s.Sections = append(s.Sections, sect)
	}
	return s, nil
}

func readSectionFromDir(dir string) (section, error) {
	s := section{}

	configPath := filepath.Join(dir, configFile)
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return s, err
	}
	err = json.Unmarshal(data, &s)
	if err != nil {
		return s, err
	}

	cssPath := filepath.Join(dir, cssFile)
	buffer, err := ioutil.ReadFile(cssPath)
	if err != nil {
		return s, err
	}
	s.Code = string(buffer)
	return s, err
}

// WriteToDir writes a StyleSet to an unpacked dir
func (ss StyleSet) WriteToDir(dir string) error {
	for _, s := range ss {
		cleanName := cleanNameRegexp.ReplaceAllString(s.Name, "_")
		subdir := filepath.Join(dir, cleanName)
		if err := cleanDir(subdir); err != nil {
			return err
		}
		if err := s.writeToDir(subdir); err != nil {
			return err
		}
	}
	return nil
}

func (s Style) writeToDir(dir string) error {
	for index, sect := range s.Sections {
		subdir := filepath.Join(dir, sectionDir, strconv.Itoa(index))
		if err := cleanDir(subdir); err != nil {
			return err
		}
		if err := sect.writeToDir(subdir); err != nil {
			return err
		}
	}
	s.Sections = []section{}
	configPath := filepath.Join(dir, configFile)
	buffer, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(configPath, buffer, 0600)
}

func (s section) writeToDir(dir string) error {
	cssPath := filepath.Join(dir, cssFile)
	configPath := filepath.Join(dir, configFile)

	if err := ioutil.WriteFile(cssPath, []byte(s.Code), 0600); err != nil {
		return err
	}
	s.Code = ""
	buffer, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(configPath, buffer, 0600)
}

func cleanDir(path string) error {
	if err := os.RemoveAll(path); err != nil {
		return err
	}
	return os.MkdirAll(path, 0700)
}
