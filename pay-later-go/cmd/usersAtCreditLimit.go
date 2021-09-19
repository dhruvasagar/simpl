package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// usersAtCreditLimitCmd represents the report usersAtCreditLimit command
var usersAtCreditLimitCmd = &cobra.Command{
	Use:   "users-at-credit-limit",
	Short: "Report users at credit limit",
	Long:  `Report all users who have reached their credit limit from transactions`,
	Run: func(cmd *cobra.Command, args []string) {
		users := sdb.UsersAtCreditLimit()
		if len(users) > 0 {
			for _, user := range users {
				fmt.Println(user.Name)
			}
			return
		}
		fmt.Println("No users")
	},
}

func init() {
	reportCmd.AddCommand(usersAtCreditLimitCmd)
}
