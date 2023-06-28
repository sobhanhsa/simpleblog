/*

inspiration: 
https://dribbble.com/shots/2292415-Daily-UI-001-Day-001-Sign-Up

*/
// document.cookie = "user=John; path=/; expires=Tue, 19 Jan 2038 03:14:07 GMT"

let form = document.querySelector('form');

if (document.cookie === "") {
  form.addEventListener('submit', (e) => {

    e.preventDefault();
  
    clickedbtn  = e.submitter
    
    const usernameOemail = e.target.elements[0].value;
  
    const password = e.target.elements[1].value;
    
    if (clickedbtn.innerHTML == "Sign in") {

      if ((usernameOemail === "") || (password === "")) {
        alert("please input required fields")
        return
      }

      let username = ""

      let email = ""

      if (usernameOemail.includes("@")) {
        email = usernameOemail
      } else {
        username = usernameOemail
      }

      console.log(email, username)

      fetch("http://localhost:8080/api/login", {
        method: "POST", // *GET, POST, PUT, DELETE, etc.
        mode: "cors", // no-cors, *cors, same-origin
        cache: "no-cache", // *default, no-cache, reload, force-cache, only-if-cached
        credentials: "same-origin", // include, *same-origin, omit
        headers: {
          "Content-Type": "application/json",
          // 'Content-Type': 'application/x-www-form-urlencoded',
        },
        redirect: "follow", // manual, *follow, error
        referrerPolicy: "no-referrer", // no-referrer, *no-referrer-when-downgrade, origin, origin-when-cross-origin, same-origin, strict-origin, strict-origin-when-cross-origin, unsafe-url
        body: JSON.stringify({
          Email:email,
          Username : username,
          Password : password
        }), // body data type must match "Content-Type" header
      })
      .then((res) => {
        
        console.log(res.status)
        
        if (res.status === 401) {
          console.log("incorrect password")
          
          document.getElementById("errpassword")
            .innerText = `incorrect password`

          throw `incorrect password`
        }

        if (res.status === 500 ) {
          document.getElementsByClassName("login-form").innerHTML = `<h2>something went wrong</h2>`
          throw `server error`
        }

        return res.json()
      })
      .then((data) => {
        if (data.rescode === 0) {
          document.getElementById("errusername")
            .innerText = `incorrect username or email`

          throw `incorrect email or username`
        }else if (data.user) {
          console.log(document.getElementsByClassName("errpassword"))
          console.log(document.getElementsByClassName("errusername"))
          try {
            document.getElementById("errpassword").innerText = ``
            document.getElementById("errusername").innerText = ``
            // document.getElementsByClassName("errp")[0].remove()
            // document.getElementsByClassName("errp")[1].remove()
          
          } catch {

          }
          
          document.getElementById("btns-container").innerHTML = `<p id = "userprop">
            wellcome back ${data.user.Name}
          </p>
          `
  
          var someDate = new Date();
          var numberOfDaysToAdd = 30;
          var result = someDate.setDate(someDate.getDate() + numberOfDaysToAdd);
          result = new Date(result).toUTCString()
  
          document.cookie = "user="+JSON.stringify(data.user)+"; expires="+result+"; path=./"

          setTimeout(function(){
            console.log("after 1 second")
            window.location.assign("index.html")
          }, 1500);

        } 
      })
      .catch(e => {
        console.log(e);
      })
    }  else if (clickedbtn.innerHTML == "Sign up") {
      window.location.assign("signuppage.html")
    }
  });
} else {
  
  document.getElementsByClassName("login-form")[0].innerHTML = `<h2>you are already logged in;
   <a href="/front/index.html">home</a></h2>`
   
}




//rescode (-1) === empry body
//rescode (0) === no acc founded
//rescode (1) === incorrect password
//rescode (2) === server err