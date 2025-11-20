const path = require('node:path');
const { app, BrowserWindow, ipcMain, Menu } = require('electron/main');

Menu.setApplicationMenu(null);

app.whenReady().then(() => { // don't change to .on or start breaks https://github.com/electron/electron/pull/21972
  if (process.platform !== 'linux') app.quit();
  ipcMain.handle('ping', () => 'pong');
  const win = new BrowserWindow({width: 800, height: 600,  webPreferences: { preload: path.join(__dirname, 'preload.js')} });
  win.loadFile('resources/html/index.html');
})

app.on('window-all-closed', () => { app.quit(); })
