package ook

import (
	"ook/koo"
	"os"

	"github.com/spf13/afero"
)

func (Ook *OokDir) define() interface{} {

	userHomeDir, err := os.UserHomeDir()
	koo.CheckErr(err)

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
	koo.CheckErr(err)

	afero.WriteFile(fs, Ook.lab.configfile, []byte(string(dat)), 0644)
}
