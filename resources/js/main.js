start();

const _dataFilePath = 'dataFilePath';

async function start() {
    try {
        exitIfNotLinux()
        await Neutralino.init();
        // Neutralino.events.on("trayMenuItemClicked", onTrayMenuItemClicked);
        Neutralino.events.on("windowClose", onWindowClose);
        // setTray();
        // showInfo();
        await loadDataFile();
    }
    catch (e) { console.dir(e); }
}
async function loadDataFile() {
    try {
        const path = await Neutralino.storage.getData(_dataFilePath);
        if (path) loadfile todo
    }
    catch (e) { console.dir(e); }
}
async function saveFile() {
    try {
        const path = await Neutralino.os.showSaveDialog('Save your tasks', { defaultPath: '~/periodicTasks.js' });
        await Neutralino.storage.setData(_dataFilePath, path);
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