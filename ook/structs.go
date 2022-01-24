package ook

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
