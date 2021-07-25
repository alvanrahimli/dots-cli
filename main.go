package main

import (
	"dots/commands"
	"os"
)

func main() {
	// Cut app name
	commands.DispatchCommand(os.Args[1:])
}
