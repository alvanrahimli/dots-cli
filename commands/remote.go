package commands

// Remote command adds/removes registries for package in current folder
type Remote struct {
}

func (r Remote) getArguments() []string {
	return []string{}
}

func (r Remote) checkRequirements() (bool, string) {
	return true, ""
}

func (r Remote) ExecuteCommand(opts *Opts) CommandResult {
	return CommandResult{}
}
