package cmd

import (
	"github.com/akerl/dapper/stylish"

	"github.com/spf13/cobra"
)

var packCmd = &cobra.Command{
	Use:   "pack",
	Short: "Convert directory into JSON file",
	RunE:  packRunner,
}

func init() {
	rootCmd.AddCommand(packCmd)
	packCmd.Flags().StringP("file", "f", "", "JSON file")
	packCmd.Flags().StringP("dir", "d", "", "Style directory")
}

func packRunner(cmd *cobra.Command, args []string) error {
	file, dir, err := parseArgs(cmd)
	if err != nil {
		return err
	}

	ss, err := stylish.ReadFromDir(dir)
	if err != nil {
		return err
	}

	return ss.WriteToFile(file)
}
