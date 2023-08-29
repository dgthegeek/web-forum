const btnsComent = document.querySelectorAll(".btnComent");
const comments = document.querySelectorAll(".comment1");
const btnsCancel = document.querySelectorAll(".btnCancel");
const addComment = document.querySelectorAll(".addComment");
const seeAllComments = document.querySelectorAll(".seeAllComments");
let posts = document.querySelectorAll(".post");
let cpt = 0;

Array.from(btnsCancel).forEach((btnCancel) => {
  btnCancel.addEventListener("click", (e) => {
    cpt++;
    comments = Array.from(comments).forEach((comment, idx) => {
      posts[idx].removeAttribute("hidden");
      comment.setAttribute("hidden", "hidden");
    });
  });
});

posts = Array.from(posts);
Array.from(btnsComent).forEach((btnComent, idx) => {
  btnComent.addEventListener("click", (e) => {
    if (cpt % 2 == 0) {
      comments[idx].removeAttribute("hidden");
      posts[idx].setAttribute("hidden", "hidden");
      cpt++;
    } else {
      comments[idx].setAttribute("hidden", "hidden");
      cpt++;
    }
  });
});

addComment.forEach((e) => {
  e.addEventListener("click", async function (event) {
    let comment = e.parentElement.parentElement.querySelector(".comment");
    if (!containsOnlySpacesOrNewlines(comment.value) && comment.value !== "") {
      let idPost = e.parentElement.querySelector(".idPost");
      let username = JSON.parse(localStorage.getItem("session")).username;
      event.preventDefault();

      let commentToSubmit = {
        postId: idPost.value,
        content: comment.value,
        username: username,
      };

      console.log(commentToSubmit);

      try {
        const response = await fetch("/comments/add", {
          method: "post",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(commentToSubmit),
        });

        if (response.ok) {
          alert("commment send successfuly");
          modal.style.display = "none";
          window.location.href = "/";
        }
      } catch (e) {
        alert("unable to send post");
        modal.style.display = "none";
      }
    }
  });
});

function CreateComment({ id, postID, content,username,likes,dislikes  }, displayComments) {
  let divComment = document.createElement("div");
  divComment.className = "card mb-4 bg-dark";
  let commenter = document.createElement("div")
  commenter.className="text text-light fs-4 m-3 "
  commenter.innerText=username
  let divCommentBody = document.createElement("div");
  divCommentBody.className = "card-body ";

  let divContent = document.createElement("div");
  divContent.className = "card-text pt-2 ";
  let divLikes = document.createElement("div");
  divLikes.innerHTML =
    '<button  class="connect likeButton connect btn btn-success" value="'+id+'" >Like '+likes+'</button> <button class="connect dislikeButton btn btn-danger connect ms-3 " value="'+id+'">Dislike '+dislikes+'</button>';
  divLikes.className = "justify-content-end d-flex flex-row";
  let divContentText = document.createElement("p");
  divContentText.className = "text text-light fs-5 ";
  divContentText.innerText = content;
  divContent.appendChild(divContentText);
  divCommentBody.appendChild(divContent);
  divCommentBody.appendChild(divLikes);
  divComment.appendChild(commenter)
  divComment.appendChild(divCommentBody);
  displayComments.appendChild(divComment);
}

seeAllComments.forEach((e) => {
  let cpt = false;
  e.addEventListener("click", async event => displayCommentsOnPage(event,e))
});

function containsOnlySpacesOrNewlines(str) {
  // Regular expression to match spaces and newlines
  var regex = /^[ \n]+$/;
  return regex.test(str);
}

function LikeAndDislike() {
  let allLikes = document.querySelectorAll(".likeButton");
  let allDislikes = document.querySelectorAll(".dislikeButton");
  //for likes
  allLikes.forEach((e) => {
    e.addEventListener("click", async function (event) {
      event.preventDefault();
      let idUser = JSON.parse(localStorage.getItem("session")).id;
      let idComment = Number(e.parentElement.querySelector(".likeButton").value);
      console.log(idComment)
      let data = {
        idComment,
        idUser,
      };
      console.log(data)
      try {
        const response = await fetch("/comments/likes", {
          method: "post",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(data),
        });

        if (response.ok) {
          modal.style.display = "none";
          window.location.href = "/";
        }
      } catch (e) {
        alert("unable to like comment");
        modal.style.display = "none";
      }
    });
  });

  //for dislikes
  allDislikes.forEach((e) => {
    e.addEventListener("click", async function (event) {
      event.preventDefault();
      let idUser = JSON.parse(localStorage.getItem("session")).id;
      let idComment = Number(e.parentElement.querySelector(".dislikeButton").value);
      console.log(idComment)
      let data = {
        idComment,
        idUser,
      };
      try {
        console.log("okay");
        const response = await fetch("/comments/dislikes", {
          method: "post",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(data),
        });

        if (response.ok) {
          modal.style.display = "none";
          window.location.href = "/";
        }
      } catch (e) {
        alert("unable to send post");
        modal.style.display = "none";
      }
    });
  });

}


async function displayCommentsOnPage (event,e) {
    
  let displayComments =
    e.parentElement.parentElement.querySelector(".displayComments");
  let idPost =
    e.parentElement.parentElement.parentElement.querySelector(".idPost");
  if (!cpt) {
    displayComments.removeAttribute("hidden");
    event.preventDefault();
    const response = await fetch("/comments/display")
      .then((response) => response.json())
      .then((data) => {
        // Process the retrieved data
        data.forEach((e) => {
          console.log(e)
          if (idPost.value == e.postID) {
            CreateComment(e, displayComments);
          }
        });
      });
      LikeAndDislike()
    cpt = true;
  } else {
    displayComments.innerHTML = "";
    displayComments.setAttribute("hidden", "hidden");
    cpt = false;
  }
};