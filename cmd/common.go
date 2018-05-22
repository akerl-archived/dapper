package cmd

import (
	"github.com/akerl/dapper/stylish"

	"github.com/spf13/cobra"
)

func parseArgs(cmd *cobra.Command) (string, string, error) {
	flags := cmd.Flags()

	file, err := flags.GetString("file")
	if err != nil {
		return "", "", err
	} else if file == "" {
		file = stylish.DefaultFile
	}

	dir, err := flags.GetString("dir")
	if err != nil {
		return "", "", err
	} else if dir == "" {
		dir = stylish.DefaultDir
	}

	return file, dir, nil
}
