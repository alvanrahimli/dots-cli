package commands

import (
	"dots/models"
	"fmt"
	"github.com/jessevdk/go-flags"
	"os"
)

func GetCommands() map[string]models.Command {
	return map[string]models.Command{
		"help":   Help{},
		"init":   Init{},
		"add":    Add{},
		"push":   Push{},
		"remote": Remote{},
	}
}

func DispatchCommand(args []string) {
	// Parse CLI arguments
	var opts models.Opts
	args, err := flags.ParseArgs(&opts, args)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		os.Exit(1)
	}
	opts.Arguments = args

	commands := GetCommands()

	// Print help and exit
	if commands[args[0]] == nil {
		commands["help"].ExecuteCommand(&opts)
		os.Exit(0)
	}

	// Handle command
	result := commands[args[0]].ExecuteCommand(&opts)
	if result.Code == 0 {
		fmt.Printf("Command executed successfully:\n\t %s\n", result.Message)
		os.Exit(0)
	} else {
		fmt.Printf("Error occured: %d : %s\n", result.Code, result.Message)
		os.Exit(1)
	}
}
