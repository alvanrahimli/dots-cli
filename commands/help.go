package commands

import "fmt"

// Help command prints help message
type Help struct {

}

func (h Help) getArguments() []string {
	return []string {}
}

func (h Help) checkRequirements() bool {
	return true
}

func (h Help) ExecuteCommand() {
	fmt.Println()
	fmt.Println("   Yes, this is help message")
}