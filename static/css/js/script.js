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


let searchInput = document.querySelector('#employeeList');
let tableRows = document.querySelector('#employee-table').querySelectorAll('.table-content')

searchInput.addEventListener('input', (e) => {
    let searchInputValue = e.target.value.toLowerCase();
  tableRows.forEach(row => {
    let doesRowMatch = row.textContent.toLowerCase().includes(searchInputValue);
    if (doesRowMatch) {
      row.style.display = 'table-row';
    } else {
      row.style.display = 'none';
    }
  })
})


// let search_Input = document.querySelector('#employeeTimesheet');
// let table_Rows = document.querySelector('#timesheet').querySelectorAll('.table-content')

// search_Input.addEventListener('input', (e) => {
//     let searchInputValue1 = e.target.value.toLowerCase();
//   tableRows1.forEach(row => {
//     let doesRowMatch1 = row.textContent.toLowerCase().includes(searchInputValue1);
//     if (doesRowMatch1) {
//       row.style.display = 'table-row';
//     } else {
//       row.style.display = 'none';
//     }
//   })
// })

// function employeeList() {
//   // Declare variables
//   var input, filter, table, tr, td, i, txtValue;
//   input = document.getElementById("employeeList");
//   filter = input.value.toUpperCase();
//   table = document.getElementById("employee-table");
//   tr = table.getElementsByTagName("tr");

//   // Loop through all table rows, and hide those who don't match the search query
//   for (i = 0; i < tr.length; i++) {
//     td = tr[i].getElementsByTagName("td")[2];
//     if (td) {
//       txtValue = td.textContent || td.innerText;
//       if (txtValue.toUpperCase().indexOf(filter) > -1) {
//         tr[i].style.display = "";
//       } else {
//         tr[i].style.display = "none";
//       }
//     }
//   }
// }