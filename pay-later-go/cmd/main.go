package cmd

import (
	"strings"

	"github.com/dhruvasagar/pay-later-go/db"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pay-later-go",
	Short: "A Simple Pay-Later Service",
	Long: `As a pay later service we allow our users to buy goods from a merchant
	now, and then allow them to pay for those goods at a later date`,
	SilenceUsage:  true,
	SilenceErrors: true,
	ValidArgs:     []string{"new", "update", "payback", "report"},
	Args:          cobra.ExactValidArgs(1),
}

// sdb is a reference to our DB instance
var sdb = db.New()

// Execute executes a command by name
func Execute(cmdName string) {
	args := strings.Fields(cmdName)
	rootCmd.ResetFlags()
	rootCmd.SetArgs(args)

	// Work around to reset sub command flags
	cmd, _, _ := rootCmd.Find(args)
	if cmd != nil {
		cmd.ResetFlags()
	}
	err := rootCmd.Execute()
	if err != nil {
		color.Red("%v\n", err)
		if cmd != nil {
			cmd.Help()
			return
		}
		rootCmd.Help()
	}
}
