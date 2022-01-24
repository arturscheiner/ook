/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"ook/koo"
	"os"

	git "github.com/go-git/go-git/v5"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

// bootstrapCmd represents the bootstrap command
var bootstrapCmd = &cobra.Command{
	Use:   "bootstrap",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		rstatus, _ := cmd.Flags().GetBool("redo")
		if rstatus {
			del_ook()
			get_ook()
		} else {
			get_ook()
			check_dependencies()
		}

	},
}

func init() {
	rootCmd.AddCommand(bootstrapCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// bootstrapCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	bootstrapCmd.Flags().BoolP("redo", "r", false, "Re-bootstrap the ook needed environment")
}
func del_ook() {
	var AppFs = afero.NewOsFs()

	userdir, err := os.UserHomeDir()
	koo.CheckErr(err)

	target_dir := userdir + "/.ook"

	err = AppFs.RemoveAll(target_dir)
	koo.CheckErr(err)
}

func get_ook() {
	// Clone the given repository to the given directory
	var AppFs = afero.NewOsFs()

	userdir, err := os.UserHomeDir()
	koo.CheckErr(err)

	target_dir := userdir + "/.ook"

	tde, err := afero.DirExists(AppFs, target_dir)
	koo.CheckErr(err)

	if tde {
		fmt.Println(target_dir + " directory already exists!")
		return
	}
	_, err = git.PlainClone(target_dir, false, &git.CloneOptions{
		URL:      "https://github.com/arturscheiner/.ook.git",
		Progress: os.Stdout,
	})

	koo.CheckErr(err)
}

func check_dependencies() {
	ve := koo.CommandExists("vagrant")
	if ve {
		fmt.Println("vagrant is installed, check plugins for this system")
	}
}
