/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/zalando/go-keyring"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// setupCmd represents the setup command
var setupCmd = &cobra.Command{
	Use:   "setup [username] [password]",
	Short: "Saves Axway Open-Banking login credentials",
	Long: 
	`Uses OS Keyring to securely store your credentials.
	Only need to run this once.
	Run it again to update your credentials.`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("setup called")
		service := "amplify-ob"
		user := args[0]
		pass := args[1]

		// set username
		err := keyring.Set(service, "user", user)
		if err != nil {
			log.Fatal(err)
		}

		// set password
		err = keyring.Set(service, "pass", pass)
		if err != nil {
			log.Fatal(err)
		}

		// get username
		username, err := keyring.Get(service, "user")
		if err != nil {
			log.Fatal(err)
		}
		log.Println(username)

		// get password
		password, err := keyring.Get(service, "password")
		if err != nil {
			log.Fatal(err)
		}		
		log.Println(password)		
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setupCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
