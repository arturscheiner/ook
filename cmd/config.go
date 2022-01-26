/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"ook/ook"

	"github.com/spf13/cobra"
)

var config ook.Config

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("You must choose one of the comands")
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("install", "i", false, "Install the ook needed environment files")
	// configCmd.Flags().BoolP("redo", "r", false, "Re-install the ook needed environment")
	// configCmd.Flags().BoolP("uninstall", "u", false, "Uninstall the ook needed environment")
	// configCmd.Flags().BoolP("check", "c", false, "Check the installation and dependencies")
}
