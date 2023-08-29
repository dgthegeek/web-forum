var loginForm = document.getElementById("login-form")
var email = document.getElementById("login-email")
var password = document.getElementById("login-password")


console.log(loginForm)

loginForm.addEventListener("submit", async  e =>{
  e.preventDefault()
  var emailData = email.value
  var passwordData = password.value

  var credentials = {
    "email": emailData,
    "password":passwordData
  }

  var jsonData = JSON.stringify(credentials)

  try {
    const response = await fetch('/auth/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: jsonData
    });
  
    if (!response.ok) {
      alert("Something went wrong");
      return;
    }
  
    const data = await response.json();
    localStorage.setItem("session", JSON.stringify(data));
    localStorage.setItem("isConnected","true")
    window.location.href = '/';
  } catch (error) {
    // Handle any error that occurred during the request
    console.error(error);
  }

})