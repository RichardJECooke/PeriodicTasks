import * as _fileHelper from './fileHelper.ts';
import * as _ui from './ui.ts';

start();

async function start() {
    try {
        await Neutralino.init();
        Neutralino.events.on("windowClose", onWindowClose);
        exitIfNotLinux(); 
        (window as any)._ui = _ui; // so html can call ui
        // Neutralino.events.on("trayMenuItemClicked", onTrayMenuItemClicked);
        // setTray();
        // showInfo();
        await _fileHelper.loadPreviouslyUsedDataFile();
        await _ui.addAddTaskButton();
    }
    catch (e) { console.dir(e); }
}

// Old neutralino sample code: -----

function showInfo() {
    (document.getElementById('info') as HTMLDivElement).innerHTML = `
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

function onTrayMenuItemClicked(event: Neutralino.events.trayMenuItemClickedEvent) {
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