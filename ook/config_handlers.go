package ook

import (
	"encoding/base64"
	"fmt"
	"ook/koo"
	"os"
	"strings"

	git "github.com/go-git/go-git/v5"
	"github.com/rs/zerolog/log"
	"github.com/spf13/afero"
)

func (c *Config) handleInit() {

	c.handleConfigInstall()

	c.Fs.Mkdir(c.Dir.Lab.Root, 0755)

	dat, err := os.ReadFile(c.Dir.Home.Vagranfile)
	koo.CheckErr(err)

	scaler_sh, err := os.ReadFile(c.Dir.Home.Scaler_sh)
	koo.CheckErr(err)

	master_sh, err := os.ReadFile(c.Dir.Home.Master_sh)
	koo.CheckErr(err)

	worker_sh, err := os.ReadFile(c.Dir.Home.Worker_sh)
	koo.CheckErr(err)

	common_sh, err := os.ReadFile(c.Dir.Home.Common_sh)
	koo.CheckErr(err)

	vagrantfile := strings.Replace(string(dat), "conf.rb", c.Dir.Home.Confrb, 5)
	vagrantfile = strings.Replace(string(vagrantfile), "lab.rb", c.Dir.Home.Labrb, 5)

	scaler_sh_enc := base64.StdEncoding.EncodeToString([]byte(scaler_sh))
	master_sh_enc := base64.StdEncoding.EncodeToString([]byte(master_sh))
	worker_sh_enc := base64.StdEncoding.EncodeToString([]byte(worker_sh))
	common_sh_enc := base64.StdEncoding.EncodeToString([]byte(common_sh))

	vagrantfile = strings.Replace(string(vagrantfile), "oo-SCALER_SH-oo", string(scaler_sh_enc), 5)
	vagrantfile = strings.Replace(string(vagrantfile), "oo-MASTER_SH-oo", string(master_sh_enc), 5)
	vagrantfile = strings.Replace(string(vagrantfile), "oo-WORKER_SH-oo", string(worker_sh_enc), 5)
	vagrantfile = strings.Replace(string(vagrantfile), "oo-COMMON_SH-oo", string(common_sh_enc), 5)

	afero.WriteFile(c.Fs, "Vagrantfile", []byte(string(vagrantfile)), 0644)

	c.CreateFiles()
}

func (c *Config) CreateFiles() {
	dat, err := os.ReadFile(c.Dir.Home.Confrb)
	koo.CheckErr(err)

	afero.WriteFile(c.Fs, c.Dir.Lab.Configfile, []byte(string(dat)), 0644)
}

func (c *Config) GetVersion() string {
	dat, err := os.ReadFile(c.Dir.Home.Version)
	koo.CheckErr(err)

	return strings.TrimSuffix(string(dat), "\n")
}

func (c *Config) handleConfigInstall() {
	// Clone the given repository to the given directory
	log.Info().Msg("Handling Install")
	tde, err := afero.DirExists(c.Fs, c.Dir.Home.Root)
	koo.CheckErr(err)

	if tde {
		fmt.Println(c.Dir.Home.Root + " directory already exists!")
		return
	}
	_, err = git.PlainClone(c.Dir.Home.Root, false, &git.CloneOptions{
		URL:      "https://github.com/arturscheiner/.ook.git",
		Progress: os.Stdout,
	})

	koo.CheckErr(err)
}

func (c *Config) handleConfigRedo() {
	log.Info().Msg("Handling Redo")
	c.handleConfigUninstall()
	c.handleConfigInstall()
}

func (c *Config) handleConfigUninstall() {
	log.Info().Msg("Handling Uninstall")
	err := c.Fs.RemoveAll(c.Dir.Home.Root)
	koo.CheckErr(err)
}

func (c *Config) handleConfigInspect() {
	log.Info().Msg("Handling Check")
	ve := koo.CommandExists("vagrant")
	if ve {
		fmt.Println("vagrant is installed, check plugins for this system")
	}
}
