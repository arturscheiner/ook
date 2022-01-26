/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"ook/koo"

	"os"
	"os/exec"

	"github.com/janeczku/go-spinner"
	"github.com/spf13/cobra"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// upCmd represents the up command
var upCmd = &cobra.Command{
	Use:   "up",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		//rstdout, _ := cmd.Flags().GetBool("stdout")
		//up(rstdout)

		koo.Execute("vagrant up")

	},
}

func init() {
	rootCmd.AddCommand(upCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// upCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	upCmd.Flags().BoolP("stdout", "s", false, "Show output messages when running")
}

func up(o bool) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	cmd := exec.Command("vagrant", "up")
	if o {
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			log.Print("cmd.Run() failed with %s\n", err)
		}

	} else {
		s := spinner.StartNew("This may take some time...")
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			log.Print("cmd.Run() failed with %s\n", err)
		}
		s.Stop()
	}

	fmt.Println("Your ook lab is up and running!")

	koo.OokSsh("vagrant", "vagrant", "10.8.8.10", 22, "bash -c 'kubectl get nodes -o wide'")
}
