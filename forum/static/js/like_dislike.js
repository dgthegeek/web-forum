
var toHide = Array.from(document.querySelectorAll(".connect"))

var isConnected = localStorage.getItem('isConnected');

window.addEventListener("DOMContentLoaded",e=>{
  if (!isConnected){
    toHide.forEach(el=> el.setAttribute("hidden",true))
  }else{
    toHide.forEach(el => el.removeAttribute('hidden'))
  }
})
