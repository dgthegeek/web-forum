// Make the cookie outdated in order to delete it
document.querySelectorAll(".logout").forEach(el =>{
    el.addEventListener("submit", ()=>{
        document.cookie ="option-share-session=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;";
    })
})