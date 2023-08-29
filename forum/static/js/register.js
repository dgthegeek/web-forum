var registrationForm = document.getElementById("registration-form");
var loginUsername = document.getElementById("username")
var loginEmail = document.getElementById("email")
var loginPassword = document.getElementById("password")
var confirmLoginPassword = document.getElementById("confirmpassword")

registrationForm.addEventListener("submit", async e =>{
  e.preventDefault()
  let userData = {}
  let emailData = loginEmail.value
  let usernameData = loginUsername.value
  let passwordData = loginPassword.value
  let confirmPasswordData = confirmLoginPassword.value


  if (passwordData != confirmPasswordData){
    alert("make sure your password is the same")
  }
  userData.email = emailData
  userData.password = passwordData
  userData.username = usernameData

  let jsonData = JSON.stringify(userData)

  try{
    let response = await fetch('/auth/register', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: jsonData
    })
  
    if (!response.ok){
      alert("something went wrong")
    }
  
   // let data = await response.json()
   // localStorage.setItem("session", JSON.stringify(data))
   // localStorage.setItem("isConnected","true")
   // window.location.href="/"
    alert("user registered successfully")
    window.location.href="/login"

  }catch(e){
    alert("something went wrong")
    console.log(e)
  }
    
});

