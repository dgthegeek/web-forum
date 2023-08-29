var addPostForm = document.getElementById("addPostForm");

var postTitle = document.getElementById("post-title");
var postContent = document.getElementById("content");
var checkboxes = document.querySelectorAll('input[type="checkbox"]');


addPostForm.addEventListener("submit", async function(event) {
  event.preventDefault()

  let checked = Array.prototype.slice.call(checkboxes).some(function(checkbox) {
    return checkbox.checked;
  });
  

  const categories = Array.from(checkboxes).filter(checkbox => checkbox.checked).map(checkbox => checkbox.value)
  const author = JSON.parse(localStorage.getItem("session")).username


  let post = {
    "title" : postTitle.value,
    "content" : postContent.value,
    "categories": categories,
    "author" : author,
    "likedBy": []
  }


  if (!checked) {
    alert("Please select at least one checkbox.");
  }else if (checked){
    try{
      const response = await fetch("/posts/create",{
        method:'post',
        headers:{
          'Content-Type':'application/json'
        },
        body: JSON.stringify(post)
      });
    
      if (response.ok){
        alert("post sent successfully")
        modal.style.display = "none";
        window.location.href="/"
      }
  
    }catch(e){
      alert("unable to send post")
      modal.style.display = "none";
    }
  }
});
