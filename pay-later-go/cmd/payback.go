package cmd

import (
	"fmt"
	"strconv"

	"github.com/dhruvasagar/pay-later-go/utils"
	"github.com/spf13/cobra"
)

// paybackCmd represents the payback command
var paybackCmd = &cobra.Command{
	Use:           "payback",
	Short:         "Pay back a users dues",
	Example:       "  payback user1 300",
	SilenceUsage:  true,
	SilenceErrors: true,
	Args:          cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		userName := args[0]
		user, err := sdb.FindUserByName(userName)
		if err != nil {
			fmt.Println("Invalid User: ", userName)
			return
		}

		amount, _ := strconv.ParseFloat(args[1], 64)
		_, err = sdb.CreatePayback(user, amount)
		if err != nil {
			fmt.Println("Unable to payback:", err)
			return
		}
		fmt.Printf("%s(dues: %s)\n", user.Name, utils.FormatFloat(sdb.UserDues(user)))
	},
}

func init() {
	rootCmd.AddCommand(paybackCmd)
}
