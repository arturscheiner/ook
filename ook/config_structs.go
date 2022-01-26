package ook

import (
	"github.com/spf13/afero"
)

type Config struct {
	Fs  afero.Fs
	Dir OokDir
	Run Run
}

type OokDir struct {
	Home    OokHome
	Lab     OokLab
	Vagrant Vagrant
}

type OokHome struct {
	Root       string
	Sh         string
	Rb         string
	Vagranfile string
	Scaler_sh  string
	Master_sh  string
	Worker_sh  string
	Common_sh  string
	Confrb     string
	Labrb      string
	Version    string
}

type OokLab struct {
	Root       string
	Configfile string
	Workers    string
	Masters    string
	Hosts      string
	Ipas       string
}

type Vagrant struct {
	Root     string
	Machine  string
	Provider string
}

type Run interface {
	Init()
	Install()
}
