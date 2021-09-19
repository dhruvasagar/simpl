package cmd

import (
	"fmt"
	"strconv"

	"github.com/dhruvasagar/pay-later-go/db"
	"github.com/spf13/cobra"
)

func merchantFromArgs(args []string) db.Merchant {
	dp := args[1][:len(args[1])-1]
	discountPercentage, _ := strconv.ParseFloat(dp, 64)
	return db.Merchant{
		Name:               args[0],
		DiscountPercentage: discountPercentage,
	}
}

// merchantCmd represents the new merchant command
var merchantCreateCmd = &cobra.Command{
	Use:           "merchant",
	Short:         "Create a new merchant",
	Example:       "new merchant m1 1%",
	SilenceUsage:  true,
	SilenceErrors: true,
	Args:          cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		m, err := sdb.CreateMerchant(merchantFromArgs(args))
		if err != nil {
			fmt.Println("Unable to create merchant: ", err)
			return
		}
		fmt.Println(m)
	},
}

// merchantUpdateCmd represents the update merchant command
var merchantUpdateCmd = &cobra.Command{
	Use:           "merchant",
	Short:         "Update a merchant",
	Example:       "update merchant m1 2%",
	SilenceUsage:  true,
	SilenceErrors: true,
	Args:          cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		merchant, err := sdb.FindMerchantByName(args[0])
		if err != nil {
			fmt.Println("Invalid merchant: ", args[0])
			return
		}

		m := merchantFromArgs(args)
		m.ID = merchant.ID
		_, err = sdb.UpdateMerchant(m)
		if err != nil {
			fmt.Println("Unable to update merchant: ", err)
			return
		}
		fmt.Println(m)
	},
}

func init() {
	newCmd.AddCommand(merchantCreateCmd)
	updateCmd.AddCommand(merchantUpdateCmd)
}
