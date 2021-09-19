package cmd

import (
	"fmt"
	"strconv"

	"github.com/dhruvasagar/pay-later-go/db"
	"github.com/spf13/cobra"
)

func userFromArgs(args []string) db.User {
	creditLimit, _ := strconv.ParseFloat(args[2], 64)
	return db.User{
		Name:        args[0],
		Email:       args[1],
		CreditLimit: creditLimit,
	}
}

// userCmd represents the new user command
var userCreateCmd = &cobra.Command{
	Use:           "user",
	Short:         "Create a new user",
	Example:       "new user u1 u1@email.in 1000",
	Args:          cobra.ExactArgs(3),
	SilenceUsage:  true,
	SilenceErrors: true,
	Run: func(cmd *cobra.Command, args []string) {
		u, err := sdb.CreateUser(userFromArgs(args))
		if err != nil {
			fmt.Println("Unable to create user: ", err)
			return
		}
		fmt.Println(u)
	},
}

// userUpdateCmd represents the update user command
var userUpdateCmd = &cobra.Command{
	Use:           "user",
	Short:         "Update a user",
	Example:       "update user u1 u1@email.in 1000",
	Args:          cobra.ExactArgs(3),
	SilenceUsage:  true,
	SilenceErrors: true,
	Run: func(cmd *cobra.Command, args []string) {
		user, err := sdb.FindUserByName(args[0])
		if err != nil {
			fmt.Println("Invalid user:", args[0])
			return
		}

		u := userFromArgs(args)
		u.ID = user.ID
		_, err = sdb.UpdateUser(u)
		if err != nil {
			fmt.Println("Unable to update user: ", err)
			return
		}
		fmt.Println(u)
	},
}

func init() {
	newCmd.AddCommand(userCreateCmd)
	updateCmd.AddCommand(userUpdateCmd)
}
