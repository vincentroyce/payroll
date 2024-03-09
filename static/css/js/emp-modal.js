// Get the modal
var modal1 = document.getElementById("empModal");

// Get the button that opens the modal
var btn = document.getElementById("add-emp");

// Get the <span> element that closes the modal
var span = document.getElementsByClassName("close")[0];

// When the user clicks on the button, open the modal
btn.onclick = function() {
  modal1.style.display = "block";
}

// When the user clicks on <span> (x), close the modal
span.onclick = function() {
  modal1.style.display = "none";
}

var currentTab = 0; // Current tab is set to be the first tab (0)
showTab(currentTab); // Display the current tab

function showTab(n) {
  // This function will display the specified tab of the form ...
  var x = document.getElementsByClassName("tab");
  x[n].style.display = "block";
  // ... and fix the Previous/Next buttons:
  if (n == 0) {
    document.getElementById("prevBtn").style.display = "none";
  } else {
    document.getElementById("prevBtn").style.display = "inline";
  }
  if (n == (x.length - 1)) {
    document.getElementById("nextBtn").style.display = "none"; // Hide the next button
    
  } else {
    document.getElementById("nextBtn").innerHTML = "Next";
    document.getElementById("nextBtn").style.display = "inline"; // Show the next button
  }
  // ... and run a function that displays the correct step indicator:
  fixStepIndicator(n)
}

function nextPrev(n) {
  // This function will figure out which tab to display
  var x = document.getElementsByClassName("tab");
  // Exit the function if any field in the current tab is invalid:
  if (n == 1 && !validateForm()) return false;
  // Hide the current tab:
  x[currentTab].style.display = "none";
  // Increase or decrease the current tab by 1:
  currentTab = currentTab + n;
  // if you have reached the end of the form... :
  if (currentTab >= x.length) {
    //...the form gets submitted:
    document.getElementById("regform").submit();
    return false;
  }
  // Otherwise, display the correct tab:
  showTab(currentTab);
}

function validateForm() {
  return true; 
}

function fixStepIndicator(n) {
  // This function removes the "active" class of all steps...
  var i, x = document.getElementsByClassName("step");
  for (i = 0; i < x.length; i++) {
    x[i].className = x[i].className.replace(" active", "");
  }
  //... and adds the "active" class to the current step:
  x[n].className += " active";  
}

//edit button modal
var modal2 = document.getElementById("modal1");

var btn1 = document.querySelectorAll("#edit");

var span1 = document.getElementsByClassName("close1")[0];

btn1.forEach(function(a) {
    
  a.addEventListener('click', function() {
      
      modal2.style.display = "block";
  });
});

span1.onclick = function () {
  modal2.style.display = "none"
}

var modal3 = document.getElementById("modal2");

var span2 = document.getElementsByClassName("close2")[0];

var employeeNames = document.querySelectorAll('.employee-name');

employeeNames.forEach(function(e) {
    
    e.addEventListener('click', function() {
        
        modal3.style.display = "block";
    });
});

span2.onclick = function () {
  modal3.style.display = "none"
} 
