package cmd

import (
	"github.com/m2keloop/cli/cmd/db"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{}

func init() {
	rootCmd.AddCommand(wordCmd)
	rootCmd.AddCommand(db.GetCmd())
	rootCmd.AddCommand(timeCmd)
}

func Executor() error {
	return rootCmd.Execute()
}
