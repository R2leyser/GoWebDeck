const buttons = document.getElementsByClassName('scrip gbutton');

for (let i = 0; i < buttons.length; i++) {
    buttons.item(i).addEventListener('click', _ => {
        try {
            document.getElementsByClassName("layout").item(0).requestFullscreen();
            console.log(buttons.item(i).getAttribute("data-id"));
            const response = fetch('/scripts/' + buttons.item(i).getAttribute("data-id"), {
                method: 'post',
            });
            console.log('Completed!', response);
        } catch (err) {
            console.error(`Error: ${err}`);
        }
    });
}

const ramServidor = document.getElementById('ram-servidor');
const cpuServidor = document.getElementById('cpu-servidor');

const ramPc = document.getElementById('ram-pc');
const cpuPc = document.getElementById('cpu-pc');

const urlPC = 'http://' + window.location.hostname + ':8000'
const urlServidor = 'http://' + window.location.hostname + ':8000/servidor'

const updateTime = 1000;
var ramObj

function updateRamPc() {

    var xhr = new XMLHttpRequest();
    xhr.open('GET', urlPC + '/ram', true);

    xhr.onreadystatechange = () => {
        if (xhr.readyState === 4 && xhr.status === 200) {
            ramObj = JSON.parse(JSON.parse(xhr.response))
            ramPc.innerText = "RAM: " + ramObj.porcento + "%"
        } else {
            ramPc.innerText = "failed to connect"
        }
    };

    xhr.send();
}

updateRamPc()
setInterval(updateRamPc, updateTime);

function updateCpuPc() {
    var xhr = new XMLHttpRequest();
    xhr.open('GET', urlPC + '/cpu', true);

    xhr.onreadystatechange = () => {
        if (xhr.readyState == 4 && xhr.status === 200) {
            cpuObj = JSON.parse(JSON.parse(xhr.response))
            cpuPc.innerText = "CPU: " + cpuObj.porcento + "%"
        } else {
            cpuPc.innerText = "failed to connect"
        }
    };

    xhr.send();
}

updateCpuPc()
setInterval(updateCpuPc, updateTime);

let ramObjServidor

function updateRamServidor() {

    var xhr = new XMLHttpRequest();
    xhr.open('GET', urlServidor + '/ram', true);

    xhr.onreadystatechange = () => {
        if (xhr.readyState === 4 && xhr.status === 200) {
            ramObj = JSON.parse(JSON.parse(xhr.response))
            ramServidor.innerText = "RAM: " + ramObj.porcento + "%"
        } else {
            ramServidor.innerText = "failed to connect"
        }
    };

    xhr.send();
}

updateRamServidor()
setInterval(updateRamServidor, updateTime);

function updateCpuServidor() {
    var xhr = new XMLHttpRequest();
    xhr.open('GET', urlServidor + '/cpu', true);

    xhr.onreadystatechange = () => {
        if (xhr.readyState == 4 && xhr.status === 200) {
            cpuObj = JSON.parse(JSON.parse(xhr.response))
            cpuServidor.innerText = "CPU: " + cpuObj.porcento + "%"
        } else {
            cpuServidor.innerText = "failed to connect"
        }
    };

    xhr.send();
}

updateCpuServidor()
setInterval(updateCpuServidor, updateTime);

