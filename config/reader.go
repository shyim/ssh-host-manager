package config

import (
	"io/ioutil"
	"path"
	"regexp"
	"strconv"
	"strings"
)

type GroupCollection []*Group

var re = regexp.MustCompile(`(?m)\s*(?P<key>\w+)\s*(?P<value>\w.*)$`)

func GetConfig() (groups GroupCollection) {
	files := GetConfigFiles()

	if files == nil {
		return groups
	}

	for _, file := range files {
		data, err := ioutil.ReadFile(file)

		currentGroup := &Group{Name: path.Base(file)}
		groups = append(groups, currentGroup)

		if err != nil {
			continue
		}

		lines := strings.Split(string(data), "\n")
		var currentConfig = Config{}

		for _, line := range lines {
			if len(line) == 0 {
				continue
			}

			key, value := getLineInformation(line)

			if key == "Host" {
				if len(currentConfig.Name) > 0 {
					currentGroup.Configs = append(currentGroup.Configs, currentConfig)
				}
				currentConfig = NewConfig()
			}

			switch key {
			case "Host":
				currentConfig.Name = value
				break
			case "Hostname":
				currentConfig.Host = value
				break
			case "Port":
				currentConfig.Port, _ = strconv.ParseInt(value, 10, 64)
				break
			case "User":
				currentConfig.User = value
				break
			case "Identityfile":
				currentConfig.IdentityFile = value
				break
			case "Proxycommand":
				currentConfig.ProxyCommand = value
				break
			case "Forwardagent":
				currentConfig.ForwardAgent = value
				break
			}
		}

		if len(currentConfig.Name) > 0 {
			currentGroup.Configs = append(currentGroup.Configs, currentConfig)
		}
	}

	return groups
}

func getLineInformation(line string) (string, string) {
	match := re.FindStringSubmatch(line)
	result := make(map[string]string)
	for i, name := range re.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}

	result["key"] = strings.Title(strings.ToLower(result["key"]))

	return result["key"], result["value"]
}
