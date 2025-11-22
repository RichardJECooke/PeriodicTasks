import * as path from 'node:path';
import { fileURLToPath } from 'node:url';
import { app, BrowserWindow, ipcMain, Menu } from 'electron';
import * as _constants from '../tsCommon/1constants.ts';

const currentDir = path.dirname(fileURLToPath(import.meta.url));

if (_constants.isProduction) Menu.setApplicationMenu(null); //disables devtools and menu

app.whenReady().then(() => {
  if (process.platform !== 'linux') app.quit();
  ipcMain.handle('ping', () => 'pong');
  const mainWindow = new BrowserWindow({width: 1280, height: 768,  webPreferences: {
    preload: path.join(currentDir, 'preload.ts'),
    devTools: true
  } });
  if (!_constants.isProduction) mainWindow.webContents.openDevTools();
  mainWindow.loadFile('../../build/index.html');
})

app.on('window-all-closed', () => { app.quit(); })
