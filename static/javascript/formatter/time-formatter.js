const minute = 60 * 1000;
const hour = 60 * minute;
const day = 24 * hour;
const week = 7 * day;
const month = 30 * day;
const year = 365 * day;

function TimeFormatter(postTime) {
    const date = new Date(postTime);
    const currentTime = new Date();
    const timeDifference = currentTime - date;
    let timeAgo = "";

    if (timeDifference < minute) {
        timeAgo = "Just now";
    } else if (timeDifference < hour) {
        const minutesAgo = Math.floor(timeDifference / minute);
        timeAgo = `${minutesAgo} ${minutesAgo === 1 ? "minute" : "minutes"} ago`;
    } else if (timeDifference < day) {
        const hoursAgo = Math.floor(timeDifference / hour);
        timeAgo = `${hoursAgo} ${hoursAgo === 1 ? "hour" : "hours"} ago`;
    } else if (timeDifference < week) {
        const daysAgo = Math.floor(timeDifference / day);
        timeAgo = `${daysAgo} ${daysAgo === 1 ? "day" : "days"} ago`;
    } else if (timeDifference < month) {
        const weeksAgo = Math.floor(timeDifference / week);
        timeAgo = `${weeksAgo} ${weeksAgo === 1 ? "week" : "weeks"} ago`;
    } else if (timeDifference < year) {
        const monthsAgo = Math.floor(timeDifference / month);
        timeAgo = `${monthsAgo} ${monthsAgo === 1 ? "month" : "months"} ago`;
    } else {
        const yearsAgo = Math.floor(timeDifference / year);
        timeAgo = `${yearsAgo} ${yearsAgo === 1 ? "year" : "years"} ago`;
    }
    return timeAgo;
}

const createdAtElements = document.querySelectorAll(".createdAt");
createdAtElements.forEach(timeElement => {
    const formattedTime = TimeFormatter(timeElement.textContent);
    timeElement.textContent = formattedTime;
});
