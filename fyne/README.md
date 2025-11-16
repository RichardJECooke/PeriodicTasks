# Periodic Tasks

Each top level folder is the same app written in a different framework.

Fyne is the best one and used for release.

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
---
- make same project in fyne/go, gtk4/ts, avalonia/c#
- add open source licence file
- set project to public
- deploy to flatpak

## Run

```sh
cd ~/code/PeriodicTasks/fyne;
```

## History

- First tried Neutralino.js. Little things like file picker file extension filter didn't work. Then big things like file watcher didn't work. Gave up. Next trying Fyne and Go.