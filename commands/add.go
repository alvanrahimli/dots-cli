package commands

// Add command when name is specified adds new application to package.
type Add struct {
}

func (a Add) getArguments() []string {
	return []string{}
}

func (a Add) checkRequirements() bool {
	return true
}

func (a Add) ExecuteCommand(opts *Opts) CommandResult {
	return CommandResult{}
}
