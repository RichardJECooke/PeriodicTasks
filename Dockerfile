FROM node:23-slim
RUN apt update && apt-get install -y --no-install-recommends libgtk-3-0 libwebkit2gtk-4.0-37 libappindicator3-1 libdbus-glib-1-2 && rm -rf /var/lib/apt/lists/*
WORKDIR /app