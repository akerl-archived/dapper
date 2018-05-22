package cmd

import (
	"github.com/akerl/dapper/stylish"

	"github.com/spf13/cobra"
)

var unpackCmd = &cobra.Command{
	Use:   "unpack",
	Short: "Convert JSON file to directories",
	RunE:  unpackRunner,
}

func init() {
	rootCmd.AddCommand(unpackCmd)
	unpackCmd.Flags().StringP("file", "f", "", "JSON file")
	unpackCmd.Flags().StringP("dir", "d", "", "Style directory")
}

func unpackRunner(cmd *cobra.Command, args []string) error {
	file, dir, err := parseArgs(cmd)
	if err != nil {
		return err
	}

	ss, err := stylish.ReadFromFile(file)
	if err != nil {
		return err
	}

	return ss.WriteToDir(dir)
}
