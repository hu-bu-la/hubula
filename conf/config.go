package conf

import "time"

type Config struct {
	Title    string
	App      app
	Webstatic webstatic
	Webvar webvar
	DB       map[string]mysql `toml:"mysql"`
	Redis    map[string]redis
	Releases releases
	Company  Company
	Song    []song
}

type app struct {
	Name string
	Owner string
	Author  string
	Release time.Time
	Port int
	Org     string `toml:"organization"`
	Mark    string
}

type webstatic struct {
	Assets string
	Public string
	Favicon string
}

type webvar struct {
	Views  string
}

type mysql struct {
	DriverName string
	Host string
	Ports int
	User string
	Pwd string
	Database string
	IsRunning bool
	ConnMax  int `toml:"connection_max"`
	//Enabled bool
}

type redis struct {
	Host string
	Port int
}

type releases struct {
	Release []string
	Tags    [][]interface{}
}

type Company struct {
	Name   string
	Detail detail
}

type detail struct {
	Type string
	Addr string
	ICP  string
}

type song struct {
	Name string
	Dur  duration `toml:"duration"`
}

type duration struct {
	time.Duration
}

func (d *duration) UnmarshalText(text []byte) error {
	var err error
	d.Duration, err = time.ParseDuration(string(text))
	return err
}