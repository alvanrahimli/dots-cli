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
	Installed   bool   `long:"installed" description:"Should app print only installed apps"`
	//AppName		string `short:"a" long:"app" description:"App to add to package"`
	Arguments []string
}

func NewManifest() Manifest {
	return Manifest{
		Id:       "",
		Name:     "",
		Author:   Author{},
		Modified: true,
		Versions: make([]Version, 0),
		Apps:     make([]App, 0),
		Remotes:  make([]RemoteAddr, 0),
	}
}

type Manifest struct {
	Id       string
	Name     string
	Author   Author
	Modified bool
	Versions []Version
	Apps     []App
	Remotes  []RemoteAddr
}

type Version struct {
	Major int
	Minor int
	Patch int
}

func (v Version) ToString() string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}

func (v Version) ToFormattedString() string {
	return fmt.Sprintf("%d_%d_%d", v.Major, v.Minor, v.Patch)
}

func (v Version) IncreaseVersionNumber(major, minor, patch int) Version {
	return Version{
		Major: v.Major + major,
		Minor: v.Minor + minor,
		Patch: v.Patch + patch,
	}
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
	AuthorToken string
	Registry    string
	ConfigUrl   string
	Handlers    map[string]Handler
}

type Handler struct {
	Version    string
	ConfigRoot string
	Dotfiles   []string
}

func (manifest *Manifest) OfferNewVersion() {
	version := manifest.Versions[len(manifest.Versions)-1]
	newVersion := version.IncreaseVersionNumber(0, 0, 1)
	fmt.Printf("Enter new version number (%s -> %s): ",
		version.ToString(), newVersion.ToString())
	_, scanErr := fmt.Scanf("%d.%d.%d",
		&newVersion.Major,
		&newVersion.Minor,
		&newVersion.Patch)
	if scanErr != nil {
		manifest.Versions = append(manifest.Versions, version.IncreaseVersionNumber(0, 0, 1))
	} else {
		manifest.Versions = append(manifest.Versions, newVersion)
	}
}

func (manifest *Manifest) LastVersion() *Version {
	return &manifest.Versions[len(manifest.Versions)-1]
}

func (opts *Opts) NormalizeFlags() {
	if opts.OutputDir == "" {
		opts.OutputDir = "."
	}
}
