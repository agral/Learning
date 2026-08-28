const calendar = document.getElementById("calendar");

const isWeekend = day => {
    return day % 7 === 0 || day % 7 === 6;
}

const getWeekdayName = day => {
    const date = new Date(todayYear, todayMonth, day);
    const dayName = new Intl.DateTimeFormat("en-US", { weekday: "short" }).format(date);
}

const today = new Date();
const todayYear = today.getFullYear();
const todayMonth = today.getMonth();
for (let day=1; day <= 31; day++) {
    calendar.insertAdjacentHTML("beforeend", `<div class="day week${isWeekend(day) ? "end" : ""}">${day}</div>`);
}
