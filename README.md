# Periodic Tasks

Each top level folder is the same app written in a different framework.

## Notes

For .NET version, use flatpak https://github.com/flathub/org.freedesktop.Sdk.Extension.dotnet9
Look at https://github.com/webui-dev/webui when it's done too.

RAM use:
- Fyne       168 MB (20 MB disk)
- Avalonia   173 MB (70 MB disk)
- Electron   870 MB
- Neutralino 950 MB
- GTK

## History

- First tried Neutralino.js. Little things like file picker file extension filter didn't work. Then big things like file watcher didn't work. Gave up. Next trying Electron.
- Electron interprocess bridge for security annoyed me. The app is also RAM heavy. Trying Fyne/Go next.
