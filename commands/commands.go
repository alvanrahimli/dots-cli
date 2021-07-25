package commands

type Command interface {
	getArguments() []string
	checkRequirements() bool
	ExecuteCommand()
}

func GetCommands() map[string]Command {
	return map[string]Command {
		"help": Help{},
		"init": Init{},
		"add": Add{},
		"push": Push{},
		"remote": Remote{},
	}
}

func DispatchCommand(args []string) {
	commands := GetCommands()
	commands[args[0]].ExecuteCommand()
}

