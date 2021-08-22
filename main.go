package main

import (
	"github.com/alvanrahimli/dots-cli/commands"
	"github.com/alvanrahimli/dots-cli/dlog"
	"github.com/alvanrahimli/dots-cli/utils"
	"os"
	"strings"
)

func main() {
	dlog.Info("========| APP STARTED (cmd: %s) |========\n", strings.Join(os.Args, " "))

	config, configErr := utils.ReadConfig()
	if configErr != nil {
		dlog.PrintToStdout(true)
		dlog.Err(configErr.Error())
		return
	}

	// Dispatch command
	withoutAppName := os.Args[1:]
	commands.DispatchCommand(config, withoutAppName)

	dlog.Info("========| APP EXITED |========\n")
	return
}
