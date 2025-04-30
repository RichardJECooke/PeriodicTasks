// This is just a sample app. You can structure your Neutralinojs app code as you wish.
// This example app is written with vanilla JavaScript and HTML.
// Feel free to use any frontend framework you like :)
// See more details: https://neutralino.js.org/docs/how-to/use-a-frontend-library

/*
    Function to display information about the Neutralino app.
    This function updates the content of the 'info' element in the HTML
    with details regarding the running Neutralino application, including
    its ID, port, operating system, and version information.
*/
function showInfo() {
    document.getElementById('info').innerHTML = `
        ${NL_APPID} is running on port ${NL_PORT} inside ${NL_OS}
        <br/><br/>
        <span>server: v${NL_VERSION} . client: v${NL_CVERSION}</span>
        `;
}

/*
    Function to open the official Neutralino documentation in the default web browser.
*/
async function openDocs() {
    await Neutralino.os.open("https://neutralino.js.org/docs");
}

/*
    Function to open a tutorial video on Neutralino's official YouTube channel in the default web browser.
*/
async function openTutorial() {
    await Neutralino.os.open("https://www.youtube.com/c/CodeZri");
}

/*
    Function to set up a system tray menu with options specific to the window mode.
    This function checks if the application is running in window mode, and if so,
    it defines the tray menu items and sets up the tray accordingly.
*/
async function setTray() {
    try {
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
        await Neutralino.os.setTray(tray);
    }
    catch (e) {
        console.dir(e);
    }
}

/*
    Function to handle click events on the tray menu items.
    This function performs different actions based on the clicked item's ID,
    such as displaying version information or exiting the application.
*/
async function onTrayMenuItemClicked(event) {
    try {
        switch(event.detail.id) {
            case "VERSION":
                console.log("Case VERSION reached."); // Log when the correct case is hit
                console.log("Attempting to show message box..."); // Log before the showMessageBox call
                await Neutralino.os.showMessageBox("Version information", `Neutralinojs server: v${NL_VERSION} | Neutralinojs client: v${NL_CVERSION}`);
                console.log("Message box call completed successfully."); // Log if the await resolves
                break;
            case "QUIT":
                // Exit the application
                await Neutralino.app.exit();
                break;
        }
    }
    catch (e) {
        console.error('Error caught in onTrayMenuItemClicked:'); // Log if the catch block is hit
        console.dir(e); // Log the error object details
    }
}

/*
    Function to handle the window close event by gracefully exiting the Neutralino application.
*/
async function onWindowClose() {
    await Neutralino.app.exit();
}

try {
    Neutralino.init();
    Neutralino.events.on("trayMenuItemClicked", onTrayMenuItemClicked);
    Neutralino.events.on("windowClose", onWindowClose);
    if(NL_OS != "Darwin") { // TODO: Fix https://github.com/neutralinojs/neutralinojs/issues/615
        setTray();
    }
    showInfo();
    await Neutralino.os.showMessageBox("Version information", `Neutralinojs server: v${NL_VERSION} | Neutralinojs client: v${NL_CVERSION}`);
}
catch (e) {
    console.dir(e);
}