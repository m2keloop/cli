package cmd

import (
	"cli/cmd/db"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{}

func init() {
	rootCmd.AddCommand(wordCmd)
	rootCmd.AddCommand(db.GetCmd())
}

func Executor() error {
	return rootCmd.Execute()
}
