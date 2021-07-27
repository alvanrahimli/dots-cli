package commands

import "fmt"

// Help command prints help message
type Help struct {
}

func (h Help) getArguments() []string {
	return []string{}
}

func (h Help) checkRequirements() (bool, string) {
	return true, ""
}

func (h Help) ExecuteCommand(opts *Opts) CommandResult {
	fmt.Println()
	fmt.Println("   Yes, this is help message")
	return CommandResult{}
}
