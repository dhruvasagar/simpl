package cmd

import (
	"fmt"

	"github.com/dhruvasagar/pay-later-go/utils"
	"github.com/spf13/cobra"
)

// totalDuesCmd represents the report totalDues command
var totalDuesCmd = &cobra.Command{
	Use:   "total-dues",
	Short: "Report total dues for all users",
	Long:  `Report total due amount for all users`,
	Run: func(cmd *cobra.Command, args []string) {
		dues := sdb.UsersTotalDues()
		for _, due := range dues {
			fmt.Printf("%s: %s\n", due.Name, utils.FormatFloat(due.Due))
		}
	},
}

func init() {
	reportCmd.AddCommand(totalDuesCmd)
}
