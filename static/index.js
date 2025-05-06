const buttons = document.getElementsByClassName('script-button');

console.log('Script buttons:', buttons);
for (let i = 0; i < buttons.length; i++) {
    buttons.item(i).addEventListener('click', _ => {
      try {     
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
