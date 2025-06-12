package config

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/labstack/gommon/log"
)

type ClientAssets struct {
	BuildPath string `json:"-"`
	BaseUri   string `json:"-"`
	AppCss    string `json:"app.css"`
	AppJs     string `json:"app.js"`
	ShareCss  string `json:"share.css"`
	ShareJs   string `json:"share.js"`
}

func NewClientAssets(buildPath, baseUri string) *ClientAssets {
	return &ClientAssets{
		BuildPath: buildPath,
		BaseUri:   baseUri,
	}
}

func (a *ClientAssets) Load(fileName string) error {
	var buildPath string

	jsonFile, err := os.ReadFile(filepath.Join(buildPath, fileName))

	if err != nil {
		return err
	}

	return json.Unmarshal(jsonFile, a)
}

func (c *Config) clientAssets() *ClientAssets {
	var buildPath string
	var staticUri string

	result := NewClientAssets(buildPath, staticUri)

	if err := result.Load("assets.json"); err != nil {
		log.Errorf("frontend: cannot read assets.json")
	}

	return result
}
