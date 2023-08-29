// Get the button that opens the modal
var addPostBtn = document.getElementById("addpost");


// Get the modal
var modal = document.getElementById("myModal");

// Get the <span> element that closes the modal
var span = document.getElementsByClassName("close")[0];

var createdPostBtn = document.getElementById("fetch_post")
var likedPostBtn = document.getElementById("fetch_liked")

// When the user clicks on the button, open the modal
addPostBtn.onclick = function() {
  modal.style.display = "block";
}

// When the user clicks on <span> (x), close the modal
span.onclick = function() {
  modal.style.display = "none";
}

// When the user clicks anywhere outside of the modal, close it
window.onclick = function(event) {
  if (event.target == modal) {
    modal.style.display = "none";
  }
}


const loginBtn = document.getElementById("login-button");
const logoutBtn = document.getElementById("logout-button");

window.addEventListener("DOMContentLoaded",e =>{
  let isConnected = localStorage.getItem("isConnected")
  if (isConnected == "true"){
    loginBtn.setAttribute("hidden","true")
    addPostBtn.removeAttribute("hidden")
    createdPostBtn.removeAttribute("hidden")
    likedPostBtn.removeAttribute("hidden")
  } 
  if (!isConnected) {
    loginBtn.removeAttribute("hidden")
    logoutBtn.setAttribute("hidden","true")
    loginBtn.removeAttribute("hidden")
    logoutBtn.setAttribute("hidden","true")
    addPostBtn.setAttribute("hidden","true")
    createdPostBtn.setAttribute("hidden","true")
    likedPostBtn.setAttribute("hidden","true")
  }
})

logoutBtn.addEventListener("click", async e=>{
  e.preventDefault()

  try{
    const response = await fetch("/auth/logout")
    if (response.ok){
      localStorage.removeItem("isConnected")
      localStorage.removeItem("session")
      loginBtn.removeAttribute("hidden")
      logoutBtn.setAttribute("hidden","true")

      window.location.href="/"
    }
  }catch(e){
    console.log(e)
    alert("something went wrong")
    
  }


})