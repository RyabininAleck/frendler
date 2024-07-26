package config

type Config struct {
	Adapter      AdapterConf      `json:"adapter"`
	DB           DBConf           `json:"db"`
	Task         TaskConf         `json:"task"`
	Logger       LoggerConf       `json:"logger"`
	GoogleOAuth2 GoogleOAuth2Conf `json:"googleOAuth2"`
}

type AdapterConf struct{}

type DBConf struct {
	Path string `json:"path"`
}

type TaskConf struct {
	Interval int `json:"interval"`
}

type LoggerConf struct {
	LogLevel  string `json:"log_level"`
	LogFormat string `json:"log_format"`
}
type GoogleOAuth2Conf struct {
	RedirectURL      string   `json:"redirect_url"`
	ClientID         string   `json:"client_id"`
	ClientSecret     string   `json:"client_secret"`
	Scopes           []string `json:"scopes"`
	OauthStateString string   `json:"oauth_state_string"`
}
