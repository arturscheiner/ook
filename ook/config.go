package ook

import (
	"ook/koo"
	"os"
	"strings"

	"github.com/spf13/afero"
)

func (c *Config) Run(r string) {
	c.Define()

	switch r {
	case "init":
		c.handleInit()
	case "install":
		c.handleConfigInstall()
	case "redo":
		c.handleConfigRedo()
	case "unistall":
		c.handleConfigUninstall()
	}
}

func (c *Config) Define() *Config {

	userHomeDir, err := os.UserHomeDir()
	koo.CheckErr(err)
	c.Fs = afero.NewOsFs()
	c.Dir.Home.Root = userHomeDir + "/.ook"
	c.Dir.Home.Sh = c.Dir.Home.Root + "/lib/sh"
	c.Dir.Home.Rb = c.Dir.Home.Root + "/lib/rb"
	c.Dir.Home.Vagranfile = c.Dir.Home.Root + "/Vagrantfile"
	c.Dir.Home.Confrb = c.Dir.Home.Root + "/conf.rb"
	c.Dir.Home.Labrb = c.Dir.Home.Root + "/lib/rb/lab.rb"
	c.Dir.Home.Master_sh = c.Dir.Home.Sh + "/master.sh"
	c.Dir.Home.Worker_sh = c.Dir.Home.Sh + "/worker.sh"
	c.Dir.Home.Scaler_sh = c.Dir.Home.Sh + "/scaler.sh"
	c.Dir.Home.Common_sh = c.Dir.Home.Sh + "/common.sh"
	c.Dir.Home.Version = c.Dir.Home.Root + "/VERSION"

	c.Dir.Lab.Root = ".ook"
	c.Dir.Lab.Configfile = c.Dir.Lab.Root + "/config.env"
	c.Dir.Lab.Masters = c.Dir.Lab.Root + "/masters"
	c.Dir.Lab.Workers = c.Dir.Lab.Root + "/workers"
	c.Dir.Lab.Hosts = c.Dir.Lab.Root + "/hosts"

	return c
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
