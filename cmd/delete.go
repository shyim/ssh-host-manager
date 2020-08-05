package cmd

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/shyim/ssh-host-manager/config"
	"github.com/urfave/cli/v2"
)

func DeleteCommand(c *cli.Context) error {
	choosenText := ""
	prompt := &survey.Select{
		Message: "Choose to delete",
		Options: getOptions(),
	}
	survey.AskOne(prompt, &choosenText)

	groups := config.GetConfig()
	var group *config.Group

	for _, g := range groups {
		for i, c := range g.Configs {
			if c.Name == choosenText {
				group = g
				g.Configs[i] = g.Configs[len(g.Configs)-1]
				g.Configs = g.Configs[:len(g.Configs)-1]
			}
		}
	}

	config.UpdateConfig(group)

	return nil
}

func getOptions() (list []string) {
	for _, group := range config.GetConfig() {
		for _, c := range group.Configs {
			list = append(list, c.Name)
		}
	}

	return list
}
