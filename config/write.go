package config

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func UpdateConfig(group *Group) {
	config := generateConfig(group)

	_ = ioutil.WriteFile(GetConfigPath()+group.Name, []byte(config), 0600)

	addIncludeIfNeeded()
}

func addIncludeIfNeeded() {
	configPath := getSshConfigPath()

	if !fileExists(configPath) {
		_ = ioutil.WriteFile(configPath, []byte("Include groups/*"), 0600)
		return
	}

	bytes, err := ioutil.ReadFile(getSshConfigPath())
	config := string(bytes)

	if err != nil {
		return
	}

	if strings.Contains(config, "Include groups/*") {
		return
	}

	config += "\nInclude groups/*"

	_ = ioutil.WriteFile(configPath, []byte(config), 0600)
}

func generateConfig(group *Group) string {
	config := ""

	for _, c := range group.Configs {
		config += "Host " + c.Name + "\n"
		config += "  Hostname " + c.Host + "\n"
		config += "  Port " + strconv.FormatInt(c.Port, 10) + "\n"

		if len(c.User) > 0 {
			config += "  User " + c.User + "\n"
		}

		if len(c.IdentityFile) > 0 {
			config += "  IdentityFile " + c.IdentityFile + "\n"
		}

		if len(c.ProxyCommand) > 0 {
			config += "  ProxyCommand " + c.ProxyCommand + "\n"
		}

		if len(c.ForwardAgent) > 0 {
			config += "  ForwardAgent " + c.ForwardAgent + "\n"
		}
		config += "\n"
	}

	return config
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
