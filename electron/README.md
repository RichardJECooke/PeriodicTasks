# Periodic Tasks

Each top level folder is the same app written in a different framework.


## Todo

- use typescript
- convert my neutralino code to node/renderer separation
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
# first run
cd ~/code/PeriodicTasks/electron;
npm install;
sudo chown root:root node_modules/electron/dist/chrome-sandbox;
sudo chmod 4755 node_modules/electron/dist/chrome-sandbox;

# run
cd ~/code/PeriodicTasks/electron;
clear; rm -rf .parcel-cache/ build/; mkdir build;  mkdir build/js; cp -r resources/. build/; npx parcel watch resources/index.html --dist-dir build --public-url ./;

# new window
clear; npx electron ./resources/tsSystem/main.ts

# compile error check
clear; npx -p typescript tsc --noEmit
```

## History

- First tried Neutralino.js. Little things like file picker file extension filter didn't work. Then big things like file watcher didn't work. Gave up. Next trying Fyne and Go.