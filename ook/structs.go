package ook

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
}

type OokLab struct {
	Root       string
	Configfile string
	Workers    string
	Masters    string
	Hosts      string
}

type OokDir struct {
	Home OokHome
	Lab  OokLab
}
