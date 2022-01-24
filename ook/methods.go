package ook

import (
	"ook/koo"
	"os"

	"github.com/spf13/afero"
)

func (Ook *OokDir) Define() interface{} {

	userHomeDir, err := os.UserHomeDir()
	koo.CheckErr(err)

	Ook.Home.Root = userHomeDir + "/.ook"
	Ook.Home.Sh = Ook.Home.Root + "/lib/Sh"
	Ook.Home.Rb = Ook.Home.Root + "/lib/rb"
	Ook.Home.Vagranfile = Ook.Home.Root + "/Vagrantfile"
	Ook.Home.Confrb = Ook.Home.Root + "/conf.rb"
	Ook.Home.Labrb = Ook.Home.Root + "/lib/rb/Lab.rb"
	Ook.Home.Master_sh = Ook.Home.Sh + "/saster.sh"
	Ook.Home.Worker_sh = Ook.Home.Sh + "/worker.sh"
	Ook.Home.Scaler_sh = Ook.Home.Sh + "/scaler.sh"
	Ook.Home.Common_sh = Ook.Home.Sh + "/common.sh"

	Ook.Lab.Root = ".ook"
	Ook.Lab.Configfile = Ook.Lab.Root + "/config.env"
	Ook.Lab.Masters = Ook.Lab.Root + "/Masters"
	Ook.Lab.Workers = Ook.Lab.Root + "/Workers"
	Ook.Lab.Hosts = Ook.Lab.Root + "/hosts"

	return Ook
}

func (Ook *OokDir) CreateFiles(fs afero.Fs) {
	dat, err := os.ReadFile(Ook.Home.Confrb)
	koo.CheckErr(err)

	afero.WriteFile(fs, Ook.Lab.Configfile, []byte(string(dat)), 0644)
}
