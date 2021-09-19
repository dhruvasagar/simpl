package cmd

import (
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a user | merchant",
	Example: `  update user user1 user1@email.in 200
  update merchant m1 1.5%`,
	SilenceUsage:  true,
	SilenceErrors: true,
	ValidArgs:     []string{"user", "merchant"},
	Args:          cobra.ExactValidArgs(1),
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
