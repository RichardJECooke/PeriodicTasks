import * as path from 'node:path';
import { fileURLToPath } from 'node:url';
import { app, BrowserWindow, ipcMain, Menu } from 'electron';

const currentDir = path.dirname(fileURLToPath(import.meta.url));

Menu.setApplicationMenu(null);

app.whenReady().then(() => {
  if (process.platform !== 'linux') app.quit();
  ipcMain.handle('ping', () => 'pong');
  const win = new BrowserWindow({width: 1280, height: 768,  webPreferences: { preload: path.join(currentDir, 'preload.ts')} });
  win.loadFile('../index.html');
})

app.on('window-all-closed', () => { app.quit(); })
