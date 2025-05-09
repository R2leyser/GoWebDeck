
const buttons = document.getElementsByClassName('script-button');

console.log('Script buttons:', buttons);
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



function updateRamServidor() {
    fetch('http://localhost:8000/ram/', {method: "GET"})
        .then(response =>  response.json() )
        .then(data =>  {
            const ramServidorValue = JSON.parse(data)
            ramServidor.textContent = 'RAM: ' + ramServidorValue.percent + ' %';
        })
}

setInterval(updateRamServidor, 1000);
// ramPc.insertAdjacentText('beforeend', 'RAM: ' + navigator.deviceMemory + ' GB');
// cpuPc.insertAdjacentText('beforeend', 'CPU: ' + navigator.hardwareConcurrency + ' Cores');
