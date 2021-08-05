# dots-cli

`dots` is CLI tool to build, version and publish config file bundles.

## TODO (shared):
- [ ] Writing comprehensive documentation
- [ ] Designing landing page
- [ ] Providing screenshots/videos for demonstration

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
    - [ ] cmd: `list`(lists apps)
      - [ ] `added` (lists added apps in package)
      - [ ] `all` (lists all possible apps from config)
      - [ ] `installed` (lists possible apps which are installed on system)
    
## TODO ([dots-server](github.com/alvanrahimli/dots-server))
### Functionality:
- [x] Modelling
- [x] Login / Register API endpoints
- [x] Push package API endpoint
- [x] Get PackageArchive endpoint
- [ ] Update & Delete packages endpoint
- [ ] Enhanced models (info, settings etc.)
- [ ] Enhanced endpoints for webapp
- [ ] Security considerations
