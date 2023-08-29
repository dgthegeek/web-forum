var createDPostBtn = document.getElementById("fetch_post");
var likedPostBtn = document.getElementById("fetch_liked");

createDPostBtn.addEventListener("click", async e =>{
  e.preventDefault()
  const userId = JSON.parse(localStorage.getItem("session")).username
  window.location.href=`/posts/created_post?id=${userId}`
})


likedPostBtn.addEventListener("click", async e=>{
  e.preventDefault()
  const userId = JSON.parse(localStorage.getItem("session")).id
  window.location.href=`/posts/liked_post?id=${userId}`
})

