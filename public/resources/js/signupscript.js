
let form = document.querySelector('form');

async function fetchUser(user) {
    const response = await fetch("http://localhost:8080/api/signup", {
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
            email : user.email,
            username : user.username,
            name : user.name,
            password : user.password
        }), // body data type must match "Content-Type" header
      })

      if (response.status === 404 ) {

        return {statuscode : 404, msg : "page not found"}

      }

      const data = await response.json()

      return {statuscode : response.status, msg : data.message}
}

if (document.cookie === "") {
    form.addEventListener('submit', async(e) => {

        e.preventDefault();

        if (e.submitter.innerText === "SIGN IN") {
            window.location.assign("loginpage.html")
        } else if (e.submitter.innerText === "SIGN UP") {
            
        const email = e.target.elements["emailfield"].value

        const username = e.target.elements["usernamefield"].value

        const name = e.target.elements["namefield"].value

        const password = e.target.elements["passwordfield"].value
        
        if ((email === "") || (password === "") || (name === "") || (username === "")) {
            alert("please input required fields")
            return
        }

        const response = await fetchUser({ email : email, username : username, name : name, password : password})

            if (response.statuscode !== 201) {
                if (response.statuscode === 500 ) {
                    console.log("server error")
                    return
                }
                if (response.statuscode === 401) {
                    if (response.msg === "invalid email" ) {
                        document.getElementById("erremail")
                        .innerText = `incorrect email`
                        return
                    }
                    if (response.msg === "invalid username"){
                        document.getElementById("errusername")
                        .innerText = `incorrect username`
                        return
                    }

                }
                if (response.statuscode === 400) {
                    if (response.msg === "taken email"){
                        document.getElementById("erremail")
                        .innerText = `takenemail`
                        return
                    }
                    if (response.msg === "taken username" ) {
                        document.getElementById("errusername")
                        .innerText = `taken username`
                        return
                    }
                }
            } else {

                const userobject = { email : email, username : username, name : name, password : "0000"}
          
                document.getElementById("btns-container").innerHTML = `<p id = "userprop">
                wellcome ${userobject.name} ; your account succesfully created
                </p>
                `

                //making expire time
                var someDate = new Date();
                var numberOfDaysToAdd = 30;
                var result = someDate.setDate(someDate.getDate() + numberOfDaysToAdd);
                result = new Date(result).toUTCString()
        
                
                //creating cookie
                document.cookie = "user="+JSON.stringify(userobject)+"; expires="+result+"; path=./"

                setTimeout(function(){
                    console.log("after 1 second")
                    window.location.assign("/")
                }, 1500);

            }
        }
    });
} else {

    document.getElementsByClassName("signup-form")[0].innerHTML = `<h2>you are already logged in;
    <a href="/">home</a></h2>`
    
}
