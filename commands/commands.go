package commands

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"os"
)

func GetCommands() map[string]Command {
	return map[string]Command{
		"help":   Help{},
		"init":   Init{},
		"add":    Add{},
		"push":   Push{},
		"remote": Remote{},
	}
}

func DispatchCommand(args []string) {
	// Parse CLI arguments
	var opts Opts
	args, err := flags.ParseArgs(&opts, args)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		os.Exit(1)
	}

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
