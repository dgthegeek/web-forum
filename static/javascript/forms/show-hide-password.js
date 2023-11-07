// this is a immediately invoked function that hide and show password
(function (){
    var isHidden = false;
document.querySelector(".show-password").addEventListener("click",(e)=>{
    document.querySelector(".show-password img").src= isHidden ? "/static/assets/eye.svg" : "/static/assets/eye-slash.svg"
isHidden = !isHidden
e.currentTarget.nextElementSibling.type= isHidden ? "text" : "password";
})
document.querySelector(".google-auth").addEventListener("click",(e)=>{
    e.preventDefault()
})

 })()

