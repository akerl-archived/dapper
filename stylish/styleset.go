package stylish

import (
	"path/filepath"
)

// StyleSet is a slice of Styles
type StyleSet []Style

// ReadFromFile reads a StyleSet from a packed file
func ReadFromFile(file string) (StyleSet, error) {
	ss := StyleSet{}
	err := readJSONFile(ss, file)
	return ss, err
}

// WriteToFile writes a StyleSet to a packed file
func (ss StyleSet) WriteToFile(file string) error {
	return writeJSONFile(ss, file)
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
