/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"amplify-ob/apis/login"
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/zalando/go-keyring"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Logs in using your stored credentials",
	Long: `
	Use the setup command to store your credentials`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("login called")
		service := "amplify-ob"
		user, err := keyring.Get(service, "user")
		if err != nil {
			log.Fatal(err)
		}
		pass, err := keyring.Get(service, "pass")
		if err != nil {
			log.Fatal(err)
		}
		login.LoginHandler(user, pass)
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
