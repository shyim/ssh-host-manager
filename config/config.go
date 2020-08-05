package config

import (
	"io/ioutil"
	"os"
	"os/user"
)

func getSshConfigPath() string {
	homeDir, _ := os.UserHomeDir()

	return homeDir + "/.ssh/config"
}

func GetConfigPath() string {
	homeDir, _ := os.UserHomeDir()
	configFolder := homeDir + "/.ssh/groups/"

	return configFolder
}

func GetConfigFiles() []string {
	configFolder := GetConfigPath()

	list, err := ioutil.ReadDir(configFolder)
	var files []string

	if err != nil {
		return files
	}

	for _, file := range list {
		files = append(files, configFolder+file.Name())
	}

	return files
}

type Group struct {
	Name    string
	Configs []Config
}

type Config struct {
	Name         string
	Host         string
	Port         int64
	User         string
	IdentityFile string
	ProxyCommand string
	ForwardAgent string
}

func NewConfig() Config {
	u, _ := user.Current()

	return Config{
		User: u.Username,
		Port: 22,
	}
}
