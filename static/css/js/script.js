let profile = document.querySelector('.profile')
let modal = document.querySelector('.modal')
let modalContent = document.querySelector('.modal-content')
let modalEdit = document.querySelector('.modal-content-edit')
let edit = document.querySelector('.edit-profile')
let tableContent = document.querySelectorAll('.table-content')
let checkBox = document.querySelectorAll('.select')
let checkAll = document.querySelector('.selectAll')

profile.addEventListener('click', () => {
  modal.style.display = 'flex';
})

edit.addEventListener('click', () => {
  modalContent.style.display = 'none';
  modal.style.justifyContent = 'center';
  modal.style.alignItems = 'center';
  modalEdit.style.display = 'flex'
})

window.onclick = function (e) {
  if (e.target == modal) {
    modal.style.display = 'none';
    modalContent.style.display = 'flex';
    modal.style.justifyContent = 'flex-end';
    modal.style.alignItems = 'flex-start';
    modalEdit.style.display = 'none'
  }
}


let loadFile = function (event) {
  let picture = document.querySelector('#main-profile')
  picture.src = URL.createObjectURL(event.target.files[0]);
};


checkBox.forEach((tabCont, index) => {
  if (tabCont.checked) {
    tableContent[index].classList.add("active");
  }
  tabCont.addEventListener("click", () => {
      tableContent[index].classList.toggle("active");

  });
});

checkAll.addEventListener("click", () => {
  if (checkAll.checked){
    checkBox.forEach(v => {
     v.checked = true;
    })
    tableContent.forEach(v => {
      v.classList.add("active");
    })
  } else {
    checkBox.forEach(v => {
      v.checked = false;
    })
    tableContent.forEach(v => {
      v.classList.remove("active");
    })
  }
})