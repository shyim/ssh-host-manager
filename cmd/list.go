package cmd

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli/v2"
	"github.com/shyim/ssh-host-manager/config"
	"os"
	"strconv"
	"strings"
)

func ListCommand(c *cli.Context) error  {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Hostname", "Port", "User", "Forward Agent?"})

	for _, g := range config.GetConfig() {
		fmt.Println(fmt.Sprintf("=== %s ===", g.Name))

		for _, c := range g.Configs {
			table.Append([]string{
				c.Name,
				c.Host,
				strconv.FormatInt(c.Port, 10),
				c.User,
				ForwardAgentText(c.ForwardAgent)})
		}
		table.Render()
		table.ClearRows()

		fmt.Println("")
	}

	return nil
}

func ForwardAgentText(str string) string  {
	str = strings.ToLower(str)

	if str == "yes" {
		return "[X]"
	}

	return "[ ]"
}
