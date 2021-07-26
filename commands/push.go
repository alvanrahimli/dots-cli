package commands

// Push command pushes package to already added registry
// Registries can be added using `dots remote add` command
type Push struct {
}

func (p Push) getArguments() []string {
	return []string{}
}

func (p Push) checkRequirements() bool {
	return true
}

func (p Push) ExecuteCommand(opts *Opts) CommandResult {
	return CommandResult{}
}
