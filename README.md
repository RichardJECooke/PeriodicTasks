# PeriodicTaskTracker

## Todo

- add open source licence file

## Run

```sh
docker run --init  -it --rm --name "app" -v ".:/app" -w "/app" node:23-alpine3.19 sh -c  "npm install && npx neu run"

docker run --init  -it --rm --name "app" -v ".:/app" -w "/app" node:23-alpine3.19 sh -c  "npm install @neutralinojs/neu "
docker run --init  -it --rm --name "app" -v ".:/app" -w "/app" node:23-alpine3.19 sh -c  "npx neu create myapp"

```