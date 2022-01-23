/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/base64"
	"os"
	"strings"

	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

type OokHome struct {
	root       string
	sh         string
	rb         string
	vagranfile string
	scaler_sh  string
	master_sh  string
	worker_sh  string
	common_sh  string
	confrb     string
	labrb      string
}

type OokLab struct {
	root       string
	configfile string
	workers    string
	masters    string
	hosts      string
}

type OokDir struct {
	home OokHome
	lab  OokLab
}

func (Ook *OokDir) define() interface{} {

	userHomeDir, err := os.UserHomeDir()
	check(err)

	Ook.home.root = userHomeDir + "/.ook"
	Ook.home.sh = Ook.home.root + "/lib/sh"
	Ook.home.rb = Ook.home.root + "/lib/rb"
	Ook.home.vagranfile = Ook.home.root + "/Vagrantfile"
	Ook.home.confrb = Ook.home.root + "/conf.rb"
	Ook.home.labrb = Ook.home.root + "/lib/rb/lab.rb"
	Ook.home.master_sh = Ook.home.sh + "/master.sh"
	Ook.home.worker_sh = Ook.home.sh + "/worker.sh"
	Ook.home.scaler_sh = Ook.home.sh + "/scaler.sh"
	Ook.home.common_sh = Ook.home.sh + "/common.sh"

	Ook.lab.root = ".ook"
	Ook.lab.configfile = Ook.lab.root + "/config.env"
	Ook.lab.masters = Ook.lab.root + "/masters"
	Ook.lab.workers = Ook.lab.root + "/workers"
	Ook.lab.hosts = Ook.lab.root + "/hosts"

	return Ook
}

func (Ook *OokDir) createFiles(fs afero.Fs) {
	dat, err := os.ReadFile(Ook.home.confrb)
	check(err)

	afero.WriteFile(fs, Ook.lab.configfile, []byte(string(dat)), 0644)
}

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

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func init_strap() {
	appfs := afero.NewOsFs()

	ook := &OokDir{}

	ook.define()

	appfs.Mkdir(ook.lab.root, 0755)

	dat, err := os.ReadFile(ook.home.vagranfile)
	check(err)

	scaler_sh, err := os.ReadFile(ook.home.scaler_sh)
	check(err)

	master_sh, err := os.ReadFile(ook.home.master_sh)
	check(err)

	worker_sh, err := os.ReadFile(ook.home.worker_sh)
	check(err)

	common_sh, err := os.ReadFile(ook.home.common_sh)
	check(err)

	vagrantfile := strings.Replace(string(dat), "conf.rb", ook.home.confrb, 5)
	vagrantfile = strings.Replace(string(vagrantfile), "lab.rb", ook.home.labrb, 5)

	scaler_sh_enc := base64.StdEncoding.EncodeToString([]byte(scaler_sh))
	master_sh_enc := base64.StdEncoding.EncodeToString([]byte(master_sh))
	worker_sh_enc := base64.StdEncoding.EncodeToString([]byte(worker_sh))
	common_sh_enc := base64.StdEncoding.EncodeToString([]byte(common_sh))

	vagrantfile = strings.Replace(string(vagrantfile), "oo-SCALER_SH-oo", string(scaler_sh_enc), 5)
	vagrantfile = strings.Replace(string(vagrantfile), "oo-MASTER_SH-oo", string(master_sh_enc), 5)
	vagrantfile = strings.Replace(string(vagrantfile), "oo-WORKER_SH-oo", string(worker_sh_enc), 5)
	vagrantfile = strings.Replace(string(vagrantfile), "oo-COMMON_SH-oo", string(common_sh_enc), 5)

	afero.WriteFile(appfs, "Vagrantfile", []byte(string(vagrantfile)), 0644)

	//create_config_env(appfs)
	ook.createFiles(appfs)
}

//func create_config_env(fs afero.Fs, ookdir string) {
//	dat, err := os.ReadFile(userdir + "/.ook/kvlab.conf.rb")
//	check(err)
//
//	afero.WriteFile(fs, ".ook/config.env", []byte(string(vagrantfile)), 0644)
//
//}
