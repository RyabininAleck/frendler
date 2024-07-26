package config

import (
	"encoding/json"
	"flag"
	"io"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	GoogleOauth      = &oauth2.Config{}
	OauthStateString = ""
)

func Get() *Config {
	configPath := flag.String("config", "./config.json", "path to the configuration file")
	flag.Parse()

	config, err := LoadConfig(*configPath)
	if err != nil {
		panic(err)
	}

	GoogleOauth = InitOAuth2Config(config.GoogleOAuth2)
	OauthStateString = config.GoogleOAuth2.OauthStateString
	return config
}

func LoadConfig(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := json.Unmarshal(bytes, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

func InitOAuth2Config(cfg GoogleOAuth2Conf) *oauth2.Config {
	return &oauth2.Config{
		RedirectURL:  cfg.RedirectURL,
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		Scopes:       cfg.Scopes,
		Endpoint:     google.Endpoint,
	}
}
