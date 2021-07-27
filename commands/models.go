package commands

type Command interface {
	getArguments() []string
	checkRequirements() (bool, string)
	ExecuteCommand(opts *Opts) CommandResult
}

type CommandResult struct {
	Code    int
	Message string
}

type Opts struct {
	OutputDir   string `short:"o" long:"output" description:"Output directory to initialize package in"`
	PackageName string `short:"n" long:"name" description:"Package name"`
	Version     string `long:"version" description:"Package version"`
	AuthorName  string `long:"author-name" description:"Author name"`
	AuthorEmail string `long:"author-email" description:"Author email"`
}

func NewManifest() Manifest {
	return Manifest{
		Id:      "",
		Version: Version{},
		Name:    "",
		Author:  Author{},
		Apps:    make([]App, 0),
		Remotes: make([]Remote, 0),
	}
}

type Manifest struct {
	Id      string
	Version Version
	Name    string
	Author  Author
	Apps    []App
	Remotes []Remote
}

type Version struct {
	Major int
	Minor int
	Patch int
}

type Author struct {
	Name  string
	Email string
}

type App struct {
	Name    string
	Version string
}

type RemoteAddr struct {
	Name string
	Url  string
}
