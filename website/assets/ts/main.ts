class Child {
    name: string;
    birthday: Date;

    constructor(name: string, birthday: Date) {
        this.name = name;
        this.birthday = birthday;
    }

    toString() {
        return this.name;
    }
}

class Week {
    start: Date;
    end: Date;

    constructor(start: Date, end: Date) {
        this.start = start;
        this.end = end;
    }
}

const start = new Date(2026, 0, 1);
const end = new Date(2026, 2, 1);
const date = new Date(start.getTime());
const childs = [
    new Child("Anna", new Date(2000, 0, 12)),
    new Child("Ayden", new Date(2000, 0, 12)),
    new Child("Ichem", new Date(2000, 0, 15)),
    new Child("Imen", new Date(2000, 0, 1)),
    new Child("Roxanne", new Date(2000, 0, 22)),
];

distributePluchie(start, childs);
// Récupérer les childs.length semaines suivantes
// Regarder si des enfants ont un anniversaire pendant ces semaines
// Si oui :
    // Essayer de donner la peluche le plus proche possible de son anniversaire
    // Pour chaque semaine :
        // Donner la peluche à l'enfant dont l'anniversaire est le plus proche de cette semaine
        // Si deux enfants sont nés le même jour et/ou dans la même semaine
            // Prendre un enfant de manière random
// Si non :
    // Distribuer les peluches de manière équitable entre chaque enfant


function distributePluchie(start: Date, childs: Child[]) {
    if(childs.length <= 0) {
        throw new Error('');
    }

    const childsPool = [...childs];

    const startWeek = getWeek(start);
    const endWeek = getWeek(new Date(start.getTime() + childs.length * daysInMilliseconds(7)));
    const date = new Date(startWeek.start.getTime());

    let i = 0;
    while (date.getTime() < endWeek.start.getTime()) {
        const week = getWeek(date);

        let closest: number = 0;
        let closestGap = timeBetweenIgnoringYear(date, childsPool[closest].birthday)
        const length = childsPool.length;
        for (let i = length - 1; i > 0; i--) {
            const child = childsPool[i];
            const isBirthdayInWeek = isBetweenIgnoringYear(child.birthday, week.start, week.end)

            if (!isBirthdayInWeek) {
                continue;
            }

            const gap = timeBetweenIgnoringYear(date, child.birthday);

            if (gap <= closestGap) {
                closestGap = gap;
                closest = i;
            }
        }

        console.log(childsPool[closest], week);

        let last = childsPool[childsPool.length - 1];
        childsPool[childsPool.length - 1] = childsPool[closest];
        childsPool[closest] = last;

        childsPool.pop();

        i++;
        date.setTime(date.getTime() + daysInMilliseconds(7));
    }
}

function timeBetweenIgnoringYear(base: Date, compared: Date): number {
    const sameYear = new Date(base.getFullYear(), compared.getMonth(), compared.getDate());
    const nextYear = new Date(base.getFullYear() + 1, compared.getMonth(), compared.getDate());

    const sameYearGap = Math.abs(sameYear.getTime() - base.getTime());
    const nextYearGap = Math.abs(nextYear.getTime() - base.getTime());

    if(sameYearGap > nextYearGap) {
        return nextYearGap;
    }

    return sameYearGap;
}

function isBetweenIgnoringYear(date: Date, start: Date, end: Date): boolean {
    const d1 = new Date(start.getFullYear(), date.getMonth(), date.getDate());
    const d2 = new Date(end.getFullYear(), date.getMonth(), date.getDate());

    return (
        (d1 >= start && d1 <= end) ||
        (d2 >= start && d2 <= end)
    );
}

function isBetween(date: Date, start: Date, end: Date): boolean {
    return start.getTime() < date.getTime() && date.getTime() < end.getTime();
}

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

function getWeek(date: Date, startDay: number = 0): Week {
    const dayInMonth = date.getDate();
    const dayInWeek = date.getDay() - startDay;

    const startTimestamp = date.getTime() - daysInMilliseconds(dayInWeek);
    const endTimestamp = date.getTime() + daysInMilliseconds(7 - dayInWeek);

    const start = new Date(startTimestamp);
    const end = new Date(endTimestamp);

    return new Week(start, end);
}

function daysInMilliseconds(days: number): number {
    return days * 24 * 60 * 60 * 1000;
}
