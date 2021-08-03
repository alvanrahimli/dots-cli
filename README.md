# dots-cli

`dots` is CLI tool to build, version and publish config file bundles.

## TODO (client):
### Functionality:
- [ ] Commands:
    - [x] cmd: `init` (initializes empty package)
    - [x] cmd: `add` (adds app to package)
    - [x] cmd: `remove` (removes app from package)
    - [x] cmd: `pack` (saves package's current state)
    - [ ] cmd: `revert` (loads previous version)
    - [ ] cmd: `remote` (adds/removes registries)
    - [ ] cmd: `push` (pushes package to added registry)
    - [ ] cmd: `install` (installs packages to appropriate directories)


- [x] __CRITIC__: Use config file for app handlers.
- [x] Some actual helpful help output
- [x] Check for existing package in current folder
- [x] Multiple version support
- [x] Modified indicator
- [x] Keep version archives in `$PACK/.vers/`


## TODO (server)
### Functionality:
- [ ] Whole planning and implementing
