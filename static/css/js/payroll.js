let searchInput = document.querySelector('#payroll');
let tableRows = document.querySelector('.pr-table').querySelectorAll('.payroll-content')

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