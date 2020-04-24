package config

import (
	"os"
	"os/user"
)

func getSshConfigPath() string {
	homeDir, _ := os.UserHomeDir()

	return homeDir + "/.ssh/config"
}

func GetConfigPath() string {
	homeDir, _ := os.UserHomeDir()

	return homeDir + "/.ssh/manager_hosts"
}

type Group struct {
	Name 		string
	Configs     []Config
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
		User:         u.Username,
	}
}