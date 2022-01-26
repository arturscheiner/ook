package ook

import (
	"encoding/base64"
	"ook/koo"
	"os"
	"strings"

	"github.com/spf13/afero"
)

func (c *Config) handleInit() {

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

	c.CreateFiles(c.Fs)
}

func (c *Config) CreateFiles(fs afero.Fs) {
	dat, err := os.ReadFile(c.Dir.Home.Confrb)
	koo.CheckErr(err)

	afero.WriteFile(fs, c.Dir.Lab.Configfile, []byte(string(dat)), 0644)
}

func (c *Config) GetVersion() string {
	dat, err := os.ReadFile(c.Dir.Home.Version)
	koo.CheckErr(err)

	return strings.TrimSuffix(string(dat), "\n")
}
