/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"ook/koo"
	"ook/ook"
	"time"

	"github.com/briandowns/spinner"
	"github.com/spf13/cobra"
)

// Initialize variable for down command
var down ook.Exec

// downCmd represents the down command
var downCmd = &cobra.Command{
	Use:   "down",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		//koo.TestOut()
		down.Run("vagrant halt")
	},
}

func init() {
	rootCmd.AddCommand(downCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// downCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// downCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func do_down() {
	s := spinner.New(spinner.CharSets[2], 100*time.Millisecond)
	s.Suffix = " This may take some time..."
	s.Start()

	//cmd.Stdout = os.Stdout
	//time.Sleep(1000 * time.Millisecond)
	//out, err := exec.Command("vagrant", "halt").Output()
	//koo.CheckErr(err)
	koo.Execute("vagrant halt")
	//fmt.Println(string(out))
	//s := spinner.StartNew("This may take some time...")
	//cmd.Stderr = os.Stderr
	//err := cmd.Run()
	//if err != nil {
	//	log.Fatalf("cmd.Run() failed with %s\n", err)
	//}
	//s.Stop()
	//c <- "Your ook lab is down!"
	//fmt.Println(c)
	//fmt.Println("Your ook lab is down!")
	s.Stop()
}
