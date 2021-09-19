package cmd

import (
	"fmt"

	"github.com/dhruvasagar/pay-later-go/utils"
	"github.com/spf13/cobra"
)

// duesCmd represents the report dues command
var duesCmd = &cobra.Command{
	Use:     "dues",
	Short:   "Report total dues for a user",
	Long:    `Report total dues pending for a user from all transactions`,
	Example: "report dues user1",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		userName := args[0]
		user, err := sdb.FindUserByName(userName)
		if err != nil {
			fmt.Println("Invalid User: ", userName)
			return
		}

		fmt.Printf("%s\n", utils.FormatFloat(sdb.UserDues(user)))
	},
}

func init() {
	reportCmd.AddCommand(duesCmd)
}
