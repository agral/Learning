const calendar = document.getElementById("calendar");

for (let day=1; day <= 31; day++) {
    calendar.insertAdjacentHTML("beforeend", `<div class="day">${day}</div>`);
}
