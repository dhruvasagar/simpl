package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// txnCmd represents the new txn command
var txnCmd = &cobra.Command{
	Use:           "txn",
	Short:         "Create a new transaction",
	Example:       "new txn user1 m1 500",
	SilenceUsage:  true,
	SilenceErrors: true,
	Args:          cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		userName := args[0]
		user, err := sdb.FindUserByName(userName)
		if err != nil {
			fmt.Println("Invalid User: ", userName)
			return
		}

		merchantName := args[1]
		merchant, err := sdb.FindMerchantByName(merchantName)
		if err != nil {
			fmt.Println("Invalid Merchant: ", merchantName)
			return
		}

		amount, _ := strconv.ParseFloat(args[2], 64)
		_, err = sdb.CreateTransaction(user, merchant, amount)
		if err != nil {
			fmt.Printf("rejected! (reason: %s)\n", err)
			return
		}
		fmt.Println("success!")
	},
}

func init() {
	newCmd.AddCommand(txnCmd)
}
