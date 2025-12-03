const start = new Date();
const end = new Date(2026, 7, 1);
const date = new Date(start.getTime());
const childs = ["Anna", "Ayden", "Ichem"];

const weeksEvents = [];
let i = 0;
while(date.getTime() < end.getTime()) {
    // Trouver la semaine en cours
    const week = getWeek(date, 1);

    // Attribuer un enfant à cette semaine
    weeksEvents.push({
        week: week,
        child: childs[i % childs.length],
    });

    // Avancer la date d'une semaine
    const nextWeekTimestamp = date.getTime() + daysInMilliseconds(7);
    date.setTime(nextWeekTimestamp);
    i++;
}

console.log(weeksEvents);

function getWeek(date: Date, startDay: number = 0) {
    const dayInMonth = date.getDate();
    const dayInWeek = date.getDay() - startDay;

    const startTimestamp = date.getTime() - daysInMilliseconds(dayInWeek);
    const endTimestamp = date.getTime() + daysInMilliseconds(7 - dayInWeek);

    const start = new Date(startTimestamp);
    const end = new Date(endTimestamp);

    return {
        start: start,
        end: end,
    };
}

function daysInMilliseconds(days: number): number {
    return days * 24 * 60 * 60 * 1000;
}
