package main

import (
	"github.com/dots/commands"
	"github.com/dots/utils"
	"os"
)

func main() {
	config := utils.ReadConfig()
	// Cut app name
	commands.DispatchCommand(config, os.Args[1:])
}
