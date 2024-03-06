let months = ['January', 'February', 'March', 'April', 'May', 'June', 'July', 'August', 'September', 'October', 'November', 'December']
let day = ['Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday', 'Sunday']
let dateText = document.querySelector('.date');
let time = new Date();
let hour = time.getHours();
let minute = time.getMinutes();
let seconds = time.getSeconds();
if (hour > 12) {
    hour -= 12;
}
if (minute < 10) {
    minute = `0${time.getMinutes()}` 
}
if (hour < 10) {
    hour = `0${hour}` 
}
if (seconds < 10) {
    seconds = `0${time.getSeconds()}`
}
if (time.getHours > 11) {
    dateText.textContent =`0${time.getDate()} ${months[time.getMonth()]}, ${time.getFullYear()} ${hour}:${minute}:${seconds}AM`;
} else {
    dateText.textContent =`0${time.getDate()} ${months[time.getMonth()]}, ${time.getFullYear()} ${hour}:${minute}:${seconds}PM`;
}
setInterval(function() {
    time = new Date();
    hour = time.getHours();
    minute = time.getMinutes();
    seconds = time.getSeconds();
    if (hour > 12) {
        hour -= 12;
    }
    if (minute < 10) {
        minute = `0${time.getMinutes()}` 
    }
    if (hour < 10) {
        hour = `0${hour}` 
    }
    if (seconds < 10) {
        seconds = `0${time.getSeconds()}`
    }
    if (time.getHours > 11) {
        dateText.textContent =`0${time.getDate()} ${months[time.getMonth()]}, ${time.getFullYear()} ${hour}:${minute}:${seconds}AM`;
    } else {
        dateText.textContent =`0${time.getDate()} ${months[time.getMonth()]}, ${time.getFullYear()} ${hour}:${minute}:${seconds}PM`;
    }
},1000)


