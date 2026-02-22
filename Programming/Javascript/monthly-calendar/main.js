const calendar = document.getElementById("calendar");

const isWeekend = day => {
    return day % 7 === 0 || day % 7 === 6;
}

for (let day=1; day <= 31; day++) {
    calendar.insertAdjacentHTML("beforeend", `<div class="day week${isWeekend(day) ? "end" : ""}">${day}</div>`);
}
