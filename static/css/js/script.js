let profile = document.querySelector('.profile')
let modal = document.querySelector('.modal')

profile.addEventListener('click', () => {
  modal.style.display = 'flex';
})

window.onclick = function (e) {
  if (e.target == modal) {
    modal.style.display = 'none';
  }
}