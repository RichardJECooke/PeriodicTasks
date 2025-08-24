# PeriodicTaskTracker

## Todo

- mark task as done today and save to task history
- mark task as done at any date
- display task history
- edit and delete task history
- add archiving and deletion of tasks
- automatic saving on change with toast
- display tasks next due date and rank by due date
- save config file and open it
- use parcel file exclude plugin - use https://www.npmjs.com/package/parcel-resolver-ignore instead
- add open source licence file
- set project to public
- make same project in gtk4/c and avalonia/c# and gtk4/zig

## Run

Run in top level: periodicTaskTracker folder

### With Node directly on host - use this

```sh
cd ~/code/periodicTaskTracker;
npm install;
npx neu update;

rm -rf .parcel-cache/ build/; mkdir build;  mkdir build/js; cp -r resources/. build/; npx parcel watch resources/js/main.ts --dist-dir build --public-url ./;  # can't build the neutralino.js file with parcel or it breaks - todo - use https://www.npmjs.com/package/parcel-resolver-ignore instead

# new terminal:
npx neu run;
# npx neu build # for later when deploying
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