package cmd

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/shyim/ssh-host-manager/config"
	"github.com/urfave/cli/v2"
	"strconv"
)

func EditCommand(c *cli.Context) error {
	choosenText := ""
	prompt := &survey.Select{
		Message: "Choose to edit",
		Options: getOptions(),
	}
	survey.AskOne(prompt, &choosenText)

	groups := config.GetConfig()
	var chooseConfig *config.Config
	gIndex := 0
	cIndex := 0

	for gI, g := range groups {
		for cI, c := range g.Configs {
			if c.Name == choosenText {
				gIndex = gI
				cIndex = cI
				chooseConfig = &c
			}
		}
	}

	if chooseConfig == nil {
		return fmt.Errorf("Cannot find chooseConfig by name %s", choosenText)
	}

	var qsEdit = []*survey.Question{
		{
			Name:     "Name",
			Prompt:   &survey.Input{Message: "Name?", Default: chooseConfig.Name},
			Validate: survey.Required,
		},
		{
			Name:     "Host",
			Prompt:   &survey.Input{Message: "Host?", Default: chooseConfig.Host},
			Validate: survey.Required,
		},
		{
			Name:   "Port",
			Prompt: &survey.Input{Message: "Port? (22)", Default: strconv.FormatInt(chooseConfig.Port, 10)},
		},
		{
			Name:   "User",
			Prompt: &survey.Input{Message: "User? (Current)", Default: chooseConfig.User},
		},
		{
			Name:   "IdentityFile",
			Prompt: &survey.Input{Message: "IdentityFile?", Default: chooseConfig.IdentityFile},
		},
		{
			Name: "ForwardAgent",
			Prompt: &survey.Confirm{
				Message: "ForwardAgent?",
				Default: chooseConfig.ForwardAgent == "yes",
			},
		},
		{
			Name:   "Jump",
			Prompt: &survey.Input{Message: "Jump-Over?"},
		},
	}

	answers := struct {
		Group        string
		Name         string
		Host         string
		Port         string
		User         string
		IdentityFile string
		ForwardAgent bool
		Jump         string
	}{}

	err := survey.Ask(qsEdit, &answers)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	port, _ := strconv.ParseInt(answers.Port, 10, 64)

	chooseConfig.Name = answers.Name
	chooseConfig.Host = answers.Host
	chooseConfig.Port = port
	chooseConfig.ForwardAgent = convertForwardAgent(answers.ForwardAgent)
	chooseConfig.User = answers.User
	chooseConfig.IdentityFile = answers.IdentityFile

	if len(answers.Jump) > 0 {
		chooseConfig.ProxyCommand = "ssh -W %h:%p " + answers.Jump
	}

	groups[gIndex].Configs[cIndex] = *chooseConfig

	config.UpdateConfig(groups[gIndex])

	return nil
}

func convertForwardAgent(agent bool) string {
	if agent {
		return "yes"
	}

	return "no"
}
