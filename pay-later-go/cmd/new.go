package cmd

import (
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new user | merchant | txn",
	Example: `  new user user1 user1@email.in 100
  new merchant m1 1.25%
  new txn user1 m1 100`,
	SilenceUsage:  true,
	SilenceErrors: true,
	ValidArgs:     []string{"user", "merchant", "txn"},
	Args:          cobra.ExactValidArgs(1),
}

func init() {
	rootCmd.AddCommand(newCmd)
}
