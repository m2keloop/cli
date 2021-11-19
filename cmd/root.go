package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{}

func init() {
	rootCmd.AddCommand(wordCmd)
}

func Executor() error {
	return rootCmd.Execute()
}
