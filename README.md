# dots

`dots` is a CLI tool to build, version and publish config file bundles.  

_Have you ever saw a screenshot of beautiful desktop customization at r/unixporn and wanted to try it out? But author did not provide dotfiles (config files of apps)! This is because he/she needs to host those files at github/gitlab, and it is not as easy as using `dots`.  
With `dots` you can build and publish your desktop rice (or any other dotfile) with only 2 commands! Installation? It also takes only 2 commands!_

# Getting Started
NO GIT REPOSITORY OR ELSE IS NEEDED  
To get started, you only need dots-cli tool. 
App does not have any dependecies.  
You can download by following [Installation guide](#installation).  
Install tool and run `dots init myfirstpack`, and you have just created your first package.

# Installation
Copy and paste following command in your terminal emulator and you are ready to go!
```
curl https://raw.githubusercontent.com/alvanrahimli/dots-cli/master/install.sh | sh
```
__You can also go to the release page and download executable (dots-cli-linux) yourself and use.__

## Usage
### Creating Package
- `dots login`                          Logs user in (essential to push package)
- `dots init <package_name>`            Initializes empty package in current directory
- `dots add <app1_name> <app2_name>`    Adds specified apps to package
- `dots add -w path/to/wallpaper.jpg`   Adds wallpaper to package
- `dots remove <app1_name>`             Removes specified apps from package
- `dots pack`                           Saves current state & ready to push
- `dots push`                           Pushes package to default registry
- `dots remote add <remote_name> <remote_address>`   Adds new remote address. 
  - For now, remote addresses should not contain trailing slash. This will be fixed in later releases
- `dots push origin`                  Pushes package to specified registry

### Installing package
- `dots get package@author.dots.rahim.li` Downloads specified package to current directory
- `dots install`                          Installs dotfiles & wallpapers
  - When installing wallpapers, app copies them to `$HOME/.local/share/backgrounds/`. You can also use them from package's folder
- `dots uninstall`                        Uninstalls dotfiles & restores previous dotfiles

___
## TODO (client):
### Functionality:
- [x] Commands:
    - [ ] cmd: `config` (tweak default registry and many more)
    - [ ] cmd: `register` (Registers new user from CLI)
    - [x] cmd: `init` (initializes empty package)
    - [x] cmd: `add` (adds app to package)
    - [x] cmd: `remove` (removes app from package)
    - [x] cmd: `pack` (saves package's current state)
    - [x] cmd: `login` (signs user in and saves token)
    - [x] cmd: `push` (pushes package to added registry)
    - [x] cmd: `remote` (adds/removes registries)
      - [x] `add` (adds remote registry)
      - [x] `remove` (removes remote registry)
    - [x] cmd: `list`(lists apps)
      - [x] `added` (lists added apps in package)
      - [x] `all` (use --installed flag to list only apps found on system) (lists all possible apps from config)
    - [x] cmd: `install` (installs packages to appropriate directories)
        - [x] Backup before installation
        - [x] Revert if installation fails
    - [x] cmd: `uninstall` (uninstalls package and returns bac .backup folder)
    - [x] cmd: `get` (downloads package)
    - [x] Find better place for logs.txt + better logging

