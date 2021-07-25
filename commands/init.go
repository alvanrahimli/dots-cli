package commands

// Init command initializes new package in given output folder
type Init struct {

}

func (i Init) getArguments() []string {
	return []string {}
}

func (i Init) checkRequirements() bool {
	return true
}

func (i Init) ExecuteCommand() {

}
