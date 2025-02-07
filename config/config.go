package config

import (
	"os"
	"path/filepath"
)

// AppConfig contains the base configuration fields required for gogetr.
type AppConfig struct {
	Name        string `long:"name" env:"NAME" default:"gogetr"`
	Debug       bool   `long:"debug" env:"DEBUG" default:"false"`
	ConfigDir   string
	RequestsDir string
}

func getDefaultConfigDir() string {
	var configFolderLocation string
	xdgConfigHome := os.Getenv("XDG_CONFIG_HOME")
	if xdgConfigHome != "" {
		configFolderLocation = xdgConfigHome
	} else {
		configFolderLocation, _ = os.UserHomeDir()
	}

	return filepath.Join(configFolderLocation, "gogetr")
}

func findOrCreateConfigDir(folder string) (string, error) {
	if folder == "" {
		folder = getDefaultConfigDir()
	}
	err := os.MkdirAll(folder, 0o755)
	if err != nil {
		return "", err
	}

	return folder, nil
}

func NewAppConfig() (*AppConfig, error) {
	// configDir, err := findOrCreateConfigDir(configDir)
	// if err != nil {
	// 	return nil, err
	// }

	appConfig := &AppConfig{
		Name:        "gogetr",
		Debug:       true,
		ConfigDir:   "",
		RequestsDir: "",
	}
	return appConfig, nil
}

// ConfigFilename returns the filename of the current config file
func (c *AppConfig) ConfigFilename() string {
	return filepath.Join(c.ConfigDir, "config.yml")
}

// RequestFilename returns the filename of the requests file
func (c *AppConfig) RequestFilename() string {
	return filepath.Join(c.ConfigDir, "requests.json")
}

