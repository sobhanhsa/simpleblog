async function publishArticle({title , category, body, hashtag}) {

    const response = await fetch(`/api/article`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          // 'Content-Type': 'application/x-www-form-urlencoded',
        },
        redirect: "follow", 
        referrerPolicy: "no-referrer", 
        body: JSON.stringify({
            Title : title,
            Category:category,
            Body :body,
            Hashtag : hashtag,
        }),
      })

    if (response.status === 400) {
        return {status : false, msg : "taken Title"}
    }

    const data = await response.json()

    return {status : true, msg : data.article}
}

function getCookie(cname) {
    let name = cname + "=";
    let decodedCookie = decodeURIComponent(document.cookie);
    let ca = decodedCookie.split(';');
    for(let i = 0; i <ca.length; i++) {
      let c = ca[i];
      while (c.charAt(0) == ' ') {
        c = c.substring(1);
      }
      if (c.indexOf(name) == 0) {
        return c.substring(name.length, c.length);
      }
    }
    return "";
}

async function checkUser(user) {
    const response = await fetch("/api/validate",{method : "GET"})

    const data = await response.json()

    if (data.user.ID !== user.ID) {
        return false
    }

    return true
}



async function main() {
    
    let form = document.querySelector('#publish-form')

    let user = getCookie("user");

    if (user === "") {
        document.getElementsByClassName("publish-form")[0].innerHTML =  `
        <h2>you are not logged in, click 
        <a href="/loginpage.html">here</a> to login</h2>
        `
        return
    }
    
    const userobj = JSON.parse(user)
    
    if (!checkUser(userobj)) {
        document.cookie = "user=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;";
        user = ""
        window.location.assign(`/`)
        return
    }


    form.addEventListener('submit', async(e) => {

        e.preventDefault();

        if (e.submitter.innerText === "CANCEL") {
            window.location.assign("/")
        } else if (e.submitter.innerText === "PUBLISH ARTICLE") {

            const title = e.target.elements["title-field"].value

            const category = e.target.elements["category"].value
    
            const body = e.target.elements["body-field"].value
    
            const hashtag = e.target.elements["hashtag-field"].value
            
            
            if ((title === "") || (body === "") ) {
                alert("please input required fields")
                return
            }
                        
            const result = await publishArticle({title, category, body, hashtag})

            if (result.status === false) {
                document.getElementById("errtitle").innerText = "taken Title"
                return
            }

            document.getElementsByClassName("publish-form")[0].innerHTML =  `
                <h2>you're article successfully created.</h2>
            `

            console.log(result.msg)

            setTimeout(function(){
                console.log("after 1 second")
                window.location.assign(`/profilepage.html?username=${result.msg.Auther}`)
            }, 1500);

        }

    })

}
main()



