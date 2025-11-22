# Periodic Tasks

https://docs.fyne.io/

## Todo

- synchronize files on save - watch file for changes and update.
  - at start/open file/save file, stop watching any other files and start watching that file
  - if that file changes, reload the tasks from it
- display errors in toasts
- export file
- import file
- show notifications when task due today
- add minimize to taskbar with quit and restore menu items
- add help/tips
- style
- write tests
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
# creation
sudo apt install golang gcc libgl1-mesa-dev xorg-dev libxkbcommon-dev;
go mod init PeriodicTaskTracker;
go get fyne.io/fyne/v2@latest;
go install fyne.io/tools/cmd/fyne@latest;


# first run
cd ~/code/PeriodicTasks/fyne;
sudo apt install golang gcc libgl1-mesa-dev xorg-dev libxkbcommon-dev;
npm install;
sudo chown root:root node_modules/electron/dist/chrome-sandbox;
sudo chmod 4755 node_modules/electron/dist/chrome-sandbox;

# run
cd ~/code/PeriodicTasks/electron;
clear; rm -rf .parcel-cache/ build/; mkdir build;  mkdir build/js; cp -r resources/. build/; npx parcel watch resources/index.html --dist-dir build --public-url ./ ;

# new window
clear; npx electron ./resources/tsSystem/main.ts

# compile error check
clear; npx -p typescript tsc --noEmit

# build
# uncomment content-security-policy in index.html
# include --no-source-maps in parcel build to remove eval which clashes with csp
# set isProduction = true in 1constants.ts to disable devtools
```

## History

- First tried Neutralino.js. Little things like file picker file extension filter didn't work. Then big things like file watcher didn't work. Gave up. Next trying Fyne and Go.