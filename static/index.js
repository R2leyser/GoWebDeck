const buttons = document.getElementsByClassName('script-button');

for (let i = 0; i < buttons.length; i++) {
    buttons.item(i).addEventListener('click', _ => {
      try {     
        document.getElementsByClassName("layout").item(0).requestFullscreen();
        console.log(buttons.item(i).getAttribute("data-id")) ;
        const response = fetch('/scripts/'+buttons.item(i).getAttribute("data-id"), {
          method: 'post',
        });
        console.log('Completed!', response);
      } catch(err) {
        console.error(`Error: ${err}`);
      }
    });
}

const ramServidor = document.getElementById('ram-servidor');
const cpuServidor = document.getElementById('cpu-servidor');

const ramPc = document.getElementById('ram-pc');
const cpuPc = document.getElementById('cpu-pc');

const urlPC = 'http://'+window.location.hostname+':8000'

var ramObj

let i = 0
function updateRamPc() {

    var xhr = new XMLHttpRequest();
    xhr.open('GET', 'http://192.168.1.102:8000/ram', true);
    
    xhr.onreadystatechange = () => {
        if (xhr.readyState === 4 && xhr.status === 200){
            ramObj = JSON.parse(JSON.parse(xhr.response))
            ramPc.innerText = "RAM: " + ramObj.porcento + "%"
        } else {
            ramPc.innerText = "failed to connect"
        }
    }; 
    
    xhr.send();
}

updateRamPc()
setInterval(updateRamPc, 1000);

function updateCpuPc() {
    var xhr = new XMLHttpRequest();
    xhr.open('GET', 'http://192.168.1.102:8000/cpu', true);
    
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
setInterval(updateCpuPc, 1000);
