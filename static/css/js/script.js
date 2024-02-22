let profile = document.querySelector('.profile')
let modal = document.querySelector('.modal')
let error = document.getElementById('error');
// let timeout = setTimeout(invalid, 3000);

profile.addEventListener('click', () => {
  modal.style.display = 'flex';
})

window.onclick = function (e) {
  if (e.target == modal) {
    modal.style.display = 'none';
  }
}

// function invalid() {
//   error.style.display = 'none';
// }

