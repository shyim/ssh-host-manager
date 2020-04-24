package cmd

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/shyim/ssh-host-manager/config"
	"github.com/urfave/cli/v2"
	"strconv"
)

var qs = []*survey.Question{
	{
		Name:     "Group",
		Prompt:   &survey.Input{Message: "Which Group?"},
		Validate: survey.Required,
	},
	{
		Name:     "Name",
		Prompt:   &survey.Input{Message: "Name?"},
		Validate: survey.Required,
	},
	{
		Name:     "Host",
		Prompt:   &survey.Input{Message: "Host?"},
		Validate: survey.Required,
	},
	{
		Name: "Port",
		Prompt: &survey.Input{Message: "Port? (22)"},
	},
	{
		Name: "User",
		Prompt:   &survey.Input{Message: "User? (Current)"},
	},
	{
		Name: "IdentityFile",
		Prompt:   &survey.Input{Message: "IdentityFile?"},
	},
	{
		Name: "ForwardAgent",
		Prompt:   &survey.Confirm{
			Message: "ForwardAgent?",
		},
	},
	{
		Name: "Jump",
		Prompt:   &survey.Input{Message: "Jump-Over?"},
	},
}

func AddCommand(c *cli.Context) error  {
	answers := struct {
		Group string
		Name string
		Host string
		Port string
		User string
		IdentityFile string
		ForwardAgent bool
		Jump string
	}{}

	err := survey.Ask(qs, &answers)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	groups := config.GetConfig()

	var group *config.Group

	for _, g := range groups {
		if g.Name == answers.Group {
			group = g
		}
	}

	if group == nil {
		group = &config.Group{Name: answers.Group}
		groups = append(groups, group)
	}

	port, err := strconv.ParseInt(answers.Port, 10, 64)

	if err != nil {
		port = 22
	}

	newConfig := config.Config{
		Name: answers.Name,
		Host: answers.Host,
		Port: port,
		ForwardAgent: convertForwardAgent(answers.ForwardAgent),
	}

	if len(answers.User) > 0 {
		newConfig.User = answers.User
	}

	if len(answers.IdentityFile) > 0 {
		newConfig.IdentityFile = answers.IdentityFile
	}

	if len(answers.Jump) > 0 {
		newConfig.ProxyCommand = "ssh -W %h:%p " + answers.Jump
	}

	group.Configs = append(group.Configs, newConfig)

	config.UpdateConfig(groups)

	return nil
}
