console.log("I'm inside")
document.querySelector(".register-form").addEventListener("submit", (e) => {
    // avoid submitting without validating datas
    // alert("ddhdd")
    e.preventDefault();
    var form = new FormData(e.target);
    let hasError = false;
    // clear error message before printing them again
    document.querySelectorAll(".error-message").forEach((err) => {
        err.textContent = "";
    });
    document.querySelectorAll("input").forEach((i) => {
        i.classList.remove("border-red-500");
    });
    // Dynamic error handling
    for (const [key, value] of form.entries()) {
        if (key !== "userID" && key !== "bio" && key !== "category") {
            if (!value) {
                console.log(key)
                hasError = true;
                const inputError = document.querySelector(`.${key}-error`);
                // outline the input that contains the error
                inputError.previousElementSibling.classList.add("border-red-500");
                // print the error message
                inputError.textContent = "*This field is required";
            }
        }
    }
    // if the user fullfilled all the requirements, datas are sent to the server
    if (!hasError) {
        e.target.submit();
    }
});
