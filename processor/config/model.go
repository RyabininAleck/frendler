package config

type Config struct {
	Adapter AdapterConf
	DB      DBConf
	Task    TaskConf
}

type AdapterConf struct{}

type DBConf struct {
	Path string
}

type TaskConf struct {
	Interval int `yaml:"Interval"`
}
