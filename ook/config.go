package ook

import (
	"ook/koo"
	"os"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/spf13/afero"
)

func (c *Config) Init() {
	log.Info().Msg("Hellow Init")
}

func (c *Config) Install() {
	log.Info().Msg("Hellow Init")
}

func (Ook *OokDir) Define() interface{} {

	userHomeDir, err := os.UserHomeDir()
	koo.CheckErr(err)

	Ook.Home.Root = userHomeDir + "/.ook"
	Ook.Home.Sh = Ook.Home.Root + "/lib/sh"
	Ook.Home.Rb = Ook.Home.Root + "/lib/rb"
	Ook.Home.Vagranfile = Ook.Home.Root + "/Vagrantfile"
	Ook.Home.Confrb = Ook.Home.Root + "/conf.rb"
	Ook.Home.Labrb = Ook.Home.Root + "/lib/rb/lab.rb"
	Ook.Home.Master_sh = Ook.Home.Sh + "/master.sh"
	Ook.Home.Worker_sh = Ook.Home.Sh + "/worker.sh"
	Ook.Home.Scaler_sh = Ook.Home.Sh + "/scaler.sh"
	Ook.Home.Common_sh = Ook.Home.Sh + "/common.sh"
	Ook.Home.Version = Ook.Home.Root + "/VERSION"

	Ook.Lab.Root = ".ook"
	Ook.Lab.Configfile = Ook.Lab.Root + "/config.env"
	Ook.Lab.Masters = Ook.Lab.Root + "/masters"
	Ook.Lab.Workers = Ook.Lab.Root + "/workers"
	Ook.Lab.Hosts = Ook.Lab.Root + "/hosts"

	return Ook
}

func (Ook *OokDir) CreateFiles(fs afero.Fs) {
	dat, err := os.ReadFile(Ook.Home.Confrb)
	koo.CheckErr(err)

	afero.WriteFile(fs, Ook.Lab.Configfile, []byte(string(dat)), 0644)
}

func (Ook *OokDir) GetVersion() string {
	dat, err := os.ReadFile(Ook.Home.Version)
	koo.CheckErr(err)

	return strings.TrimSuffix(string(dat), "\n")
}
