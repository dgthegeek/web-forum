// format the number of likes or comments.
function formatCount(count) {
    if (count < 1000) {
        return count.toString();
    } else if (count < 1000000) {
        return (count / 1000).toFixed(1) + "k";
    } else {
        return (count / 1000000).toFixed(1) + "M";
    }
}

document.querySelectorAll(".number-of-action").forEach((current) => {
    current.textContent = formatCount(current.textContent);
});

// view more / view less of na text.
document.querySelectorAll(".content").forEach((post) => {
    const content = post.querySelector(".content-text");
    const viewMore = post.querySelector(".content-view-more");

    if (content.scrollHeight > content.clientHeight) {
        viewMore.style.display = "block"; // Show the "View More" link if content is taller

        viewMore.addEventListener("click", () => {
            if (content.classList.contains("max-h-[70px]")) {
                content.classList.remove("max-h-[70px]");
                viewMore.textContent = "View Less";
            } else {
                content.classList.add("max-h-[70px]");
                viewMore.textContent = "View More";
            }
        });
    }
});
