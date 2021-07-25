package commands

import "fmt"

// Remote command adds/removes registries for package in current folder
type Remote struct {

}

func (r Remote) getArguments() []string {
	return []string {}
}

func (r Remote) checkRequirements() bool {
	return true
}

func (r Remote) ExecuteCommand() {
	fmt.Println()
	fmt.Println("   Yes, this is help message")
}
