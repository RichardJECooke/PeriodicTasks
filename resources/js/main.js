const _dataFilePathName = 'dataFilePath';
let _dataFilePath = null;
let _tasks = null;

start();

async function start() {
    try {
        await Neutralino.init();
        Neutralino.events.on("windowClose", onWindowClose);
        exitIfNotLinux();
        // Neutralino.events.on("trayMenuItemClicked", onTrayMenuItemClicked);
        // setTray();
        // showInfo();
        await loadDataFile();
        await saveDataFile();
    }
    catch (e) { console.dir(e); }
}
async function loadDataFile() {
    try {
        _dataFilePath = await Neutralino.storage.getData(_dataFilePathName);
        if (!_dataFilePath) return;
        const fileContent = await Neutralino.filesystem.readFile(_dataFilePath);
        _tasks = JSON.parse(fileContent);
    }
    catch (e) { console.dir(e); }
}
async function saveDataFile() {
    try {
        if (!_dataFilePath) chooseFileLocation();
        if (!_dataFilePath) return;
        await Neutralino.filesystem.writeFile(_dataFilePath, JSON.stringify(_tasks));
    }
    catch (e) { console.dir(e); }
}
async function chooseFileLocation() {
    try {
        const path = await Neutralino.os.showSaveDialog('Save your tasks', { defaultPath: '~/periodicTasks.js' });
        await Neutralino.storage.setData(_dataFilePathName, path);
    }
    catch (e) { console.dir(e); }
}
function showInfo() {
    document.getElementById('info').innerHTML = `
        ${NL_APPID} is running on port ${NL_PORT} inside ${NL_OS}
        <br/><br/>
        <span>server: v${NL_VERSION} . client: v${NL_CVERSION}</span>
        `;
}
function setTray() {
    if(NL_MODE != "window") {
        console.log("INFO: Tray menu is only available in the window mode.");
        return;
    }
    let tray = {
        icon: "/resources/icons/trayIcon.png",
        menuItems: [
            {id: "VERSION", text: "Get version"},
            {id: "SEP", text: "-"},
            {id: "QUIT", text: "Quit"}
        ]
    };
    Neutralino.os.setTray(tray);
}
function onTrayMenuItemClicked(event) {
    switch(event.detail.id) {
        case "VERSION":
            Neutralino.os.showMessageBox("Version information", `Neutralinojs server: v${NL_VERSION} | Neutralinojs client: v${NL_CVERSION}`);
            break;
        case "QUIT":
            Neutralino.app.exit();
            break;
    }
}
function onWindowClose() { Neutralino.app.exit(); }
function exitIfNotLinux() { if (NL_OS != 'Linux') Neutralino.app.exit(); }