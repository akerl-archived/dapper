package stylish

type section struct {
	URLs        []string `json:"urls"`
	URLPrefixes []string `json:"urlPrefixes"`
	Domains     []string `json:"domains"`
	Regexps     []string `json:"regexps"`
	Code        string   `json:"code"`
}

func readSectionFromDir(dir string) (section, error) {
	s := section{}

	err := readJSONFile(s, dir, configFile)
	if err != nil {
		return s, err
	}

	buffer, err := readTextFile(dir, cssFile)
	if err != nil {
		return s, err
	}
	s.Code = string(buffer)
	return s, err
}

func (s section) writeToDir(dir string) error {
	if err := writeTextFile([]byte(s.Code), dir, cssFile); err != nil {
		return err
	}
	s.Code = ""
	return writeJSONFile(s, dir, configFile)
}
