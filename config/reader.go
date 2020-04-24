package config

import (
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type GroupCollection []*Group

var re = regexp.MustCompile(`(?m)\s*(?P<key>\w+)\s*(?P<value>\w.*)$`)

func GetConfig() (groups GroupCollection)  {
	data, err := ioutil.ReadFile(GetConfigPath())

	if err != nil {
		return groups
	}

	lines := strings.Split(string(data), "\n")
	var currentConfig = Config{}
	var currentGroup *Group

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		key, value := getLineInformation(line)

		if key == "Group" {
			if currentGroup != nil && len(currentConfig.Name) > 0 {
				currentGroup.Configs = append(currentGroup.Configs, currentConfig)
			}

			currentGroup = nil
			currentConfig = NewConfig()
			for _, group := range groups {
				if group.Name == value {
					currentGroup = group
					break
				}
			}

			if currentGroup == nil {
				currentGroup = &Group{Name: value}
				groups = append(groups, currentGroup)
			}
		}

		if key == "Host" {
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

	if currentGroup != nil && len(currentConfig.Name) > 0 {
		currentGroup.Configs = append(currentGroup.Configs, currentConfig)
	}

	return groups
}

func getLineInformation(line string) (string, string)  {
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
