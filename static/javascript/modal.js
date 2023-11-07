document.querySelectorAll(".modal-toggler").forEach(el=>{
    el.addEventListener("click", ()=>{
        document.querySelector(".modal-wrapper").classList.remove("hidden");
    })
})
document.querySelector(".modal-closer").addEventListener("click", ()=>{
    document.querySelector(".modal-wrapper").classList.add("hidden");
})
