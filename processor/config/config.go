package config

func Get() Config {

	return Config{
		Adapter: AdapterConf{},
		DB:      DBConf{Path: "C:\\Users\\Lenovo\\GolandProjects\\frendler\\storage\\foo.dbo"},
		Task:    TaskConf{Interval: 100},
	}
}
