/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/base64"
	"ook/koo"
	"ook/ook"
	"os"
	"strings"

	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		init_strap()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

//unc check(e error) {
//	if e != nil {
//		panic(e)
//	}
//}

func init_strap() {
	appfs := afero.NewOsFs()

	ook := &ook.OokDir{}

	ook.Define()

	appfs.Mkdir(ook.Lab.Root, 0755)

	dat, err := os.ReadFile(ook.Home.Vagranfile)
	koo.CheckErr(err)

	scaler_sh, err := os.ReadFile(ook.Home.Scaler_sh)
	koo.CheckErr(err)

	master_sh, err := os.ReadFile(ook.Home.Master_sh)
	koo.CheckErr(err)

	worker_sh, err := os.ReadFile(ook.Home.Worker_sh)
	koo.CheckErr(err)

	common_sh, err := os.ReadFile(ook.Home.Common_sh)
	koo.CheckErr(err)

	vagrantfile := strings.Replace(string(dat), "conf.rb", ook.Home.Confrb, 5)
	vagrantfile = strings.Replace(string(vagrantfile), "lab.rb", ook.Home.Labrb, 5)

	scaler_sh_enc := base64.StdEncoding.EncodeToString([]byte(scaler_sh))
	master_sh_enc := base64.StdEncoding.EncodeToString([]byte(master_sh))
	worker_sh_enc := base64.StdEncoding.EncodeToString([]byte(worker_sh))
	common_sh_enc := base64.StdEncoding.EncodeToString([]byte(common_sh))

	vagrantfile = strings.Replace(string(vagrantfile), "oo-SCALER_SH-oo", string(scaler_sh_enc), 5)
	vagrantfile = strings.Replace(string(vagrantfile), "oo-MASTER_SH-oo", string(master_sh_enc), 5)
	vagrantfile = strings.Replace(string(vagrantfile), "oo-WORKER_SH-oo", string(worker_sh_enc), 5)
	vagrantfile = strings.Replace(string(vagrantfile), "oo-COMMON_SH-oo", string(common_sh_enc), 5)

	afero.WriteFile(appfs, "Vagrantfile", []byte(string(vagrantfile)), 0644)

	ook.CreateFiles(appfs)
}
