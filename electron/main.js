const { app, BrowserWindow } = require('electron/main');

const createWindow = () => {
  const win = new BrowserWindow({width: 800, height: 600 });
  win.loadFile('index.html');
}

app.whenReady().then(() => { // don't change to .on or breaks. https://github.com/electron/electron/pull/21972
  if (process.platform !== 'linux') app.quit(); // win32 linux darwin
  createWindow();
})

app.on('window-all-closed', () => {
  app.quit();
})
