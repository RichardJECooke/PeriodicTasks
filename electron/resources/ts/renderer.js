const information = document.getElementById('info');
information.innerText = `This app is using Chromium (v${versions.chrome()}), Node.js (v${versions.node()}), and Electron (v${versions.electron()})`;

async function d(){
const s = await window.versions.ping();
information.innerText = s;
}

d();