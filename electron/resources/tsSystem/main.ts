import * as path from 'node:path';
import { fileURLToPath } from 'node:url';
import { app, BrowserWindow, ipcMain, Menu } from 'electron';

const currentDir = path.dirname(fileURLToPath(import.meta.url));

Menu.setApplicationMenu(null);

app.whenReady().then(() => { // don't change to .on or start breaks https://github.com/electron/electron/pull/21972
  if (process.platform !== 'linux') app.quit();
  ipcMain.handle('ping', () => 'pong');
  const win = new BrowserWindow({width: 800, height: 600,  webPreferences: { preload: path.join(currentDir  , 'preload.ts')} });
  win.loadFile('../html/index.html');
})

app.on('window-all-closed', () => { app.quit(); })
