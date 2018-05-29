package stylish

import (
	"path/filepath"
	"strconv"
)

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

func readStyleFromDir(dir string) (Style, error) {
	s := Style{}

	err := readJSONFile(s, dir, configFile)
	if err != nil {
		return s, err
	}

	sectionGlob := filepath.Join(dir, sectionDir, "*", "style.css")
	sectionDirs, err := filepath.Glob(sectionGlob)
	if err != nil {
		return s, err
	}
	// Use index instead of name so that we count in numerical order
	for index := range sectionDirs {
		sectionDir := filepath.Join(dir, sectionDir, strconv.Itoa(index))
		sect, err := readSectionFromDir(sectionDir)
		if err != nil {
			return s, err
		}
		s.Sections = append(s.Sections, sect)
	}
	return s, nil
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
	return writeJSONFile(s, dir, configFile)
}
