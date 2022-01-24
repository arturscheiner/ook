/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"ook/koo"

	"github.com/spf13/cobra"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		check_it()
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// checkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// checkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func check_it() {

	//master, err := readFile(".kv/workers", 1)
	//if err != nil {
	//	log.Fatalf("cmd.Run() failed with %s\n", err)
	//}

	//cmd := exec.Command("vagrant", "ssh", master, "--", "kubectl get nodes -o wide")
	//s := spinner.StartNew("This may take some time...\n")
	//s.SetCharset([]string{"a", "b"})
	//cmd.Stdout = os.Stdout
	//cmd.Stderr = os.Stderr

	//stdout, _ := cmd.StdoutPipe()
	//f, _ := os.Create(".ook/stdout.log")

	//err = cmd.Run()
	//if err != nil {
	//	log.Fatalf("cmd.Run() failed with %s\n", err)
	//}

	//io.Copy(io.MultiWriter(f, os.Stdout), stdout)
	//cmd.Wait()
	//ookSsh("vagrant", "vagrant", "10.8.8.10", "/bin/bash")
	c := make(chan bool)

	go koo.Bar(-1, "executing")

	koo.OokSsh("vagrant", "vagrant", "10.8.8.10", "bash -c 'kubectl get nodes -o wide'", c)
	//
	//
	//s.Stop()
}
