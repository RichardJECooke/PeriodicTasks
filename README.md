# PeriodicTaskTracker

## Todo

- save config file and open it
- add open source licence file

## Run

Run in top level: periodicTaskTracker folder

### With Node

npm install
npx neu update
npx neu run
npx neu build

### With Docker

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