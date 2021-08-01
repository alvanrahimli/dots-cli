package models

import "fmt"

type Command interface {
	GetArguments() []string
	CheckRequirements() (bool, string)
	ExecuteCommand(opts *Opts, config *AppConfig) CommandResult
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
	//AppName		string `short:"a" long:"app" description:"App to add to package"`
	Arguments []string
}

func NewManifest() Manifest {
	return Manifest{
		Id:      "",
		Version: Version{},
		Name:    "",
		Author:  Author{},
		Apps:    make([]App, 0),
		Remotes: make([]RemoteAddr, 0),
	}
}

type Manifest struct {
	Id      string
	Version Version
	Name    string
	Author  Author
	Apps    []App
	Remotes []RemoteAddr
}

type Version struct {
	Major int
	Minor int
	Patch int
}

func (v Version) ToString() string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
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

type AppConfig struct {
	AuthorName  string
	AuthorEmail string
	Handlers    map[string]Handler
}

type Handler struct {
	Version    string
	ConfigRoot string
	Dotfiles   []string
}
