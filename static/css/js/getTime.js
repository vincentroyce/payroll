let months = ['January', 'February', 'March', 'April', 'May', 'June', 'July', 'August', 'September', 'October', 'November', 'December']
let day = ['Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday', 'Sunday']
let dateText = document.querySelector('.date');
function timeFormat() {
    let time = new Date();
    let hour = time.getHours();
    let minute = time.getMinutes();
    let seconds = time.getSeconds();
    let date = time.getDate()

    let meridiem = (hour >= 12) ? "PM":"AM";
    hour = (hour == 0 || hour == 12) ? 12 : hour % 12;
    hour = (hour < 10) ? `0${hour}`: `${hour}`;
    minute = (minute < 10) ? `0${minute}`: `${minute}`;
    seconds = (seconds < 10) ? `0${seconds}`: `${seconds}`;
    date  = (date < 10) ? `0${date}`: `${date}`;

    dateText.textContent = `${date} ${months[time.getMonth()]}, ${time.getFullYear()} ${hour}:${minute}:${seconds}${meridiem}`;
} 
timeFormat();
setInterval(timeFormat, 1000);