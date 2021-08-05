package main

import (
	"fmt"
	"github.com/alvanrahimli/dots-cli/commands"
	"github.com/alvanrahimli/dots-cli/utils"
	"os"
)

func main() {
	config, configErr := utils.ReadConfig()
	if configErr != nil {
		if os.IsNotExist(configErr) {
			// TODO: Create new config file if not exists
		}

		fmt.Printf("ERROR: %s\n", configErr.Error())
		return
	}

	// Dispatch command
	withoutAppName := os.Args[1:]
	commands.DispatchCommand(config, withoutAppName)
}
