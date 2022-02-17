package server

import (
	"errors"
	"net/url"
	"os"
)

type AppConfig struct {
	SiteUrl *url.URL
}

func NewAppConfig() (*AppConfig, error) {
	// get parameters from ENV
	rawurl, exists := os.LookupEnv("SITE_URL")
	if !exists {
		return nil, errors.New("SITE_URL was not specified")
	}

	strurl, err := url.Parse(rawurl)
	if err != nil {
		return nil, err
	}

	cfg := new(AppConfig)
	cfg.SiteUrl = strurl
	return cfg, nil
}

type Application struct {
	//HttpClient
	Config *AppConfig
}

func CreateApplication() (*Application, error) {
	cfg, err := NewAppConfig()
	if err != nil {
		return nil, err
	}

	app := new(Application)
	app.Config = cfg
	return app, nil
}
