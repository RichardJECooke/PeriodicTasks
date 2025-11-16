# Periodic Tasks

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

Run in top level: periodicTaskTracker folder

### With Node directly on host - use this

```sh
cd ~/code/PeriodicTasks/neutralino;
npm install;
npx neu update;
rm -rf .parcel-cache/ build/; mkdir build;  mkdir build/js; cp -r resources/. build/; npx parcel watch resources/index.html --dist-dir build --public-url ./;

# new terminal:
npx neu run;

# check for errors:
npx -p typescript tsc --noEmit

# build to deploy:
# npx neu build
```

### With Docker

Doesn't work properly with display integration. Better to use VM.

```sh
docker build -t neuimage .
xhost +local:docker
sudo chown -R $USER .
docker run --init -it --rm --platform linux/amd64 --name "neu" -v /tmp/.X11-unix:/tmp/.X11-unix -e DISPLAY=$DISPLAY -v ".:/app" -w "/app" neuimage sh -c "cd /app && npm install @neutralinojs/neu && npx neu update"
docker run --init -it --rm --platform linux/amd64 --name "neu" -v /tmp/.X11-unix:/tmp/.X11-unix -e DISPLAY=$DISPLAY -v ".:/app" -w "/app" neuimage sh -c "cd /app && npx neu run"
xhost -local:docker
```

## Build

```sh
docker run --init -it --rm --platform linux/amd64 --name "neu" -v /tmp/.X11-unix:/tmp/.X11-unix -e DISPLAY=$DISPLAY -v ".:/app" -w "/app" neuimage sh -c "cd /app && npx neu build"
```

## Create the app

Alpine doesn't work with neutralino. Must use node-slim

```sh
docker run --init  -it --rm --name "app" -v ".:/app" -w "/app" node:23-slim sh -c  "npx neu create myapp"
```

## Notes

- You mustn't build neutralino.js with parcel as it has hardcoded dependencies that mustn't be rewritten. The parcelIgnore plugin is used to exclude it.

## History

- First tried Neutralino.js. Little things like file picker file extension filter didn't work. Then big things like file watcher didn't work. Gave up. Next trying Fyne and Go.