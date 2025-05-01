start();


async function start() {
    try {
        exitIfNotLinux()
        await Neutralino.init();
        console.log('start');
        // Neutralino.events.on("trayMenuItemClicked", onTrayMenuItemClicked);
        Neutralino.events.on("windowClose", onWindowClose);
        // setTray();
        // showInfo();    

        await Neutralino.storage.setData('dataFilePath', JSON.stringify({ dataFilePath: '/tmp/data.js'}));
        const path = await Neutralino.storage.getData('dataFilePath');
        console.log(path);
        console.log('console end');    
    } 
    catch (e) {
        console.dir(e);
    }
}
function showInfo() {
    document.getElementById('info').innerHTML = `
        ${NL_APPID} is running on port ${NL_PORT} inside ${NL_OS}
        <br/><br/>
        <span>server: v${NL_VERSION} . client: v${NL_CVERSION}</span>
        `;
}
function openDocs() { Neutralino.os.open("https://neutralino.js.org/docs"); }
function openTutorial() { Neutralino.os.open("https://www.youtube.com/c/CodeZri"); }
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