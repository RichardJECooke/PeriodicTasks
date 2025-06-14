const _dataFilePathName = 'dataFilePath';
let _dataFilePath = null;

async function loadPreviouslyUsedDataFile() {
    try {
        _dataFilePath = await Neutralino.storage.getData(_dataFilePathName);
        if (!_dataFilePath) return;
        const fileContent = await Neutralino.filesystem.readFile(_dataFilePath);
        _tasks = JSON.parse(fileContent);
    }
    catch (e) { console.dir(e); }
}

async function openDataFile() {
    try {
        chooseOpenFileLocation();
        if (!_dataFilePath) return;
        const fileContent = await Neutralino.filesystem.readFile(_dataFilePath);
        _tasks = JSON.parse(fileContent);
    }
    catch (e) { console.dir(e); }
}

async function saveDataFile() {
    try {
        if (!_dataFilePath) chooseSaveFileLocation();
        if (!_dataFilePath) return;
        await Neutralino.filesystem.writeFile(_dataFilePath, JSON.stringify(_tasks));
    }
    catch (e) { console.dir(e); }
}
