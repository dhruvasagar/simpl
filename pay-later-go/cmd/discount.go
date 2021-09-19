package cmd

import (
	"fmt"

	"github.com/dhruvasagar/pay-later-go/utils"
	"github.com/spf13/cobra"
)

// discountCmd represents the report discount command
var discountCmd = &cobra.Command{
	Use:     "discount",
	Short:   "Report total discount for a merchant",
	Long:    `Report total discount for a merchant from all the transactions`,
	Example: "report discount m1",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		merchantName := args[0]
		merchant, err := sdb.FindMerchantByName(merchantName)
		if err != nil {
			fmt.Println("Invalid Merchant: ", merchantName)
			return
		}

		discount, err := sdb.MerchantDiscount(merchant)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%s\n", utils.FormatFloat(discount))
	},
}

func init() {
	reportCmd.AddCommand(discountCmd)
}
