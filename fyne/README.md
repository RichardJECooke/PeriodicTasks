# Periodic Tasks

https://docs.fyne.io

## Todo

- synchronize files on save - watch file for changes and update.
  - at start/open file/save file, stop watching any other files and start watching that file
  - if that file changes, reload the tasks from it
- watch config and data for changes and save if any changes
- make vue ui in fyne
- display errors in toasts
- export file
- import file
- show notifications when task due today
- add minimize to taskbar with quit and restore menu items
- add help/tips
- style
- write tests - https://docs.fyne.io/started/testing/
- add metadata - https://docs.fyne.io/started/metadata/
- minimize the build with UPX
- decide on versioning scheme and set version number
- update readme
- update neutralino.config
- bundle with flatpak, not freedesktop, use org.electronjs.Electron28.BaseApp
---
- make same project in fyne/go, gtk4/ts, avalonia/c#
- add open source licence file
- set project to public
- deploy to flatpak

## Actions

```sh
# setup go and fyne
cd ~/code/PeriodicTasks/fyne;
sudo apt install golang gcc libgl1-mesa-dev xorg-dev libxkbcommon-dev;
echo 'PATH="$PATH:$HOME/go/bin"' >> ~/.bashrc
source ~/.bashrc
# go mod init PeriodicTasks; don't run again. project already exists
# go get fyne.io/fyne/v2@latest;
# go install fyne.io/tools/cmd/fyne@latest;

# setup vs code
ctrl shift p - user settings json
add:
  "go.alternateTools": { "go": "/usr/bin/go" },
  "go.toolsGopath": "/home/me/go"

# run
cd ~/code/PeriodicTasks/fyne;
go mod tidy;
clear; go run .;

# check and test
cd ~/code/PeriodicTasks/fyne;
go fmt ./...
go vet ./...
go test ./...
go build -ldflags "-s -w" -o build/PeriodicTasks;   # remove debug and symbol data

# publish
cd ~/code/PeriodicTasks/fyne;
cd build;
fyne package -os linux -release;
cd ..
```
