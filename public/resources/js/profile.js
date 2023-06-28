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

async function fetchUserArticle(username) {
    let response
    response = await fetch(`/api/profile/${username}?limit=-1`,{method : "GET"})
    if (response.status === 400) {
        return {status : 400, msg : "no user with this username exists"}
    }
    const data = await response.json()

    console.log(data)

    return {status : 200, msg : data.message.articles}
    // const data = await response.json()
    // return data
};

async function deleteArticle(articleid) {
    let userResult = confirm("are you sure about that")

    if (userResult) {
        const response = await fetch(`/api/article/${articleid}`,{method : "DELETE"})
        if (response.status === 400) {
            alert("article was not deleted, please try again lator")
            return
        }
        location.reload()
    }

}

let user = getCookie("user");

const url = new URL(window.location.href);

const searchParams = url.searchParams;
const username = searchParams.get('username');

(async function() {

    // console.log(user)

    if (user !== "") {
        const userobj = JSON.parse(user)    
    
        //check user
        if (!checkUser(userobj)) {
            document.cookie = "user=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;";
            user = ""
        }
        
        
        if (userobj.Username !== username) {
            user = ""
        }

    }

    const resdata = await fetchUserArticle(username)
    const articles = resdata.msg

    if (user !== "" ) {

        if (articles.length === 0) {
            document.getElementById("articles").innerHTML += 
                `<h3>you don't have any articles yet</h3>
                    <p>hopefully you will start soon 😉</p>
                `
            return
        } 

        document.getElementById("articles").innerHTML = `
            <h2>your(${articles[0].Auther}) articles</h2>
        `
    
        articles.forEach(element => {
            document.getElementById("articles").innerHTML += 
            `<div class="article">
                
                <div class="content">
                    <h3 class="article-title" onclick="window.location.href='articlepage.html?articleid=${element.ID}';">
                    ${element.Title}
                    </h3>
                    <p class="category">${element.Category}</p>
                    <p class="releasedate">Date of release ${element.CreatedAt}</p>
                </div>
                <div class="article-btns">
                    <button onclick="window.location.href='updatepage.html?id=${element.ID}';" >Edit</button>
                    <button onclick="deleteArticle(${element.ID})" id="delete-btn">Delete</button>
                </div>
            </div>`
        });
    } 
    else {
        if (resdata.status === 400) {
                document.getElementById("articles").innerHTML = `
                <h2>${resdata.msg}</h2>
            `
            return
        }
        document.getElementById("articles").innerHTML = `
            <h2>${username} articles</h2>
        `

        console.log(articles.length)

        if (articles.length === 0) {
            document.getElementById("articles").innerHTML += 
                `<h3>${username} has not articles yet</h3>
                    <p>hopefully he/she will start soon</p>
                `
        } 

        articles.forEach(element => {
            document.getElementById("articles").innerHTML += 
            `<div class="article">
                <div class="content">
                    <h3 class="article-title" 
                    onclick="window.location.href='articlepage.html?articleid=${element.ID}';">${element.Title}
                    </h3>
                    <p class="category">${element.Category}</p>
                    <p class="releasedate">Date of release ${element.CreatedAt}</p>
                </div>
            </div>`
        });
        
    }
}
)();
