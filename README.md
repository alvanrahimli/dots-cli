# dots-cli documentation
___
`dots` is CLI tool to build, version and publish config file bundles.

_this document is under maintenance_

___
## TODO (client):
### Functionality:
- Some helpful help output
- ~~Check for existing package in current folder~~
- ~~Test `init` command for absolute path output dirs~~

- Commands:
    - ~~cmd: `add` (adds app to package)~~
    - cmd: `remove` (removes app from package)
    - cmd: `remote` (adds/removes registries)
    - cmd: `push` (pushes package to added registry)

### UX (Perhaps UI in future)
- Figure out a way to add new versions (maybe something like `commit`)
- Should I implement something like `revert`?
- TODO: UX : Should I ask for output folder if it is not specified as `-o / --output` flag

___
## TODO (server)
### Functionality:
- Is not planned
- Should support retrieving by manifest id (or name)

### UX (Perhaps UI in future)
- Discuss architecture
- Is `Manifest.Id` is useless? `Manifest.Name` can be used locally, when cloning remote url is enough