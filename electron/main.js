const { app, BrowserWindow } = require('electron/main');

app.whenReady().then(() => { // don't change to .on or start breaks https://github.com/electron/electron/pull/21972
  if (process.platform !== 'linux') app.quit();
  const win = new BrowserWindow({width: 800, height: 600 });
  win.loadFile('index.html');
})

app.on('window-all-closed', () => { app.quit(); })