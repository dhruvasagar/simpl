package cmd

import (
	"github.com/spf13/cobra"
)

// reportCmd represents the report command
var reportCmd = &cobra.Command{
	Use:       "report",
	Short:     "Report discount | dues | users-at-credit-limit | total-dues",
	Long:      `Generate pre-defined reports on data from the database`,
	ValidArgs: []string{"discount", "dues", "users-at-credit-limit", "total-dues"},
	Args:      cobra.ExactValidArgs(1),
}

func init() {
	rootCmd.AddCommand(reportCmd)
}
