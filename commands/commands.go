package commands

import (
	"fmt"
	"github.com/alvanrahimli/dots-cli/models"
	"github.com/jessevdk/go-flags"
	"os"
)

func GetCommands() map[string]models.Command {
	return map[string]models.Command{
		"help":   Help{},
		"init":   Init{},
		"add":    Add{},
		"remove": Remove{},
		"pack":   Pack{},
		"login":  Login{},
		"push":   Push{},
		"remote": Remote{},
		"list":   List{},
	}
}

func DispatchCommand(config *models.AppConfig, args []string) {
	// Parse CLI arguments
	var opts models.Opts
	args, err := flags.ParseArgs(&opts, args)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		os.Exit(1)
	}
	opts.Arguments = args
	opts.NormalizeFlags()

	commands := GetCommands()

	// Print help and exit
	if len(args) == 0 {
		commands["help"].ExecuteCommand(&opts, config)
		os.Exit(0)
	}
	if commands[args[0]] == nil {
		commands["help"].ExecuteCommand(&opts, config)
		os.Exit(0)
	}

	// Handle command
	result := commands[args[0]].ExecuteCommand(&opts, config)
	if result.Code == 0 {
		fmt.Printf("Command executed successfully:\n\t %s\n", result.Message)
		os.Exit(0)
	} else {
		fmt.Printf("Error occured: %s\n", result.Message)
		os.Exit(1)
	}
}
