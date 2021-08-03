# dots-cli - todo
___
`dots` is CLI tool to build, version and publish config file bundles.

_this document is under maintenance_

___
## TODO (client):
### Functionality:
- ~~__CRITIC__: Use config file for app handlers. (They all return static data, just read from file)~~
- ~~Some helpful help output~~
- ~~Check for existing package in current folder~~
- ~~Test `init` command for absolute path output dirs~~
- ~~Multiple version support~~
- ~~Modified indicator~~
- ~~Keep version archives in `$PACK/.vers/`~~


- Commands:
    - ~~cmd: `init` (initializes empty package)~~
    - ~~cmd: `add` (adds app to package)~~
    - ~~cmd: `remove` (removes app from package)~~
    - ~~cmd: `pack` (same as commit, makes package version)~~
    - cmd: `remote` (adds/removes registries)
    - cmd: `push` (pushes package to added registry)
      - DISCUSS: Should we use git?
    - cmd: `install` (installs packages to appropriate directories)
      - DISCUSS: Should we use git?

### UX (Perhaps UI in future)
- Figure out a way to add new versions (maybe something like `commit`)
- Should I implement something like `revert`?

___
## TODO (server)
### Functionality:
- Is not planned
- DISCUSS: Should we support retrieving by manifest id (or name)?

### UX (Perhaps UI in future)
- Discuss architecture
- Is `Manifest.Id` is useless? `Manifest.Name` can be used locally, when cloning remote url is enough
