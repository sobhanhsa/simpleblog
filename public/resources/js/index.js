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

async function fetchArticle(searchvalue = null) {
    let response
    if (searchvalue !== null ) {
        console.log(searchvalue)
        response = await fetch(`/api/result?search_query=${searchvalue}`,{ method :"GET"}) 
    } else {
        response = await fetch("/api/main",{ method :"GET"})
    }
    const data = await response.json()
    return data
}

async function insertArticle(articles) {

    if (typeof(articles) !== Array) {
        articles = await articles
        articles = articles.message
    }
    
    articles.forEach(element => {
        document.getElementById("articles").innerHTML += 
        `<div class="article">
            <br>
            <h3 class="article-title" 
            onclick="window.location.href='articlepage.html?articleid=${element.ID}';">${element.Title}
            </h3>
            <p class="category">${element.Category}</p>
            <p class="releasedate">Date of release ${element.CreatedAt}</p>
            <p class="auther">auther 
            <a href="/front/profilepage.html?username=${element.Auther}">${element.Auther}</a>
            </p>
            <p class="desc">${element.Body.substring(0,79)}...</p>
        </div>`
    });
}

const articles = fetchArticle()

const user = getCookie("user")

if (user !== "" ) {
    let userobj = JSON.parse(user)
    document.getElementById("user-info").innerHTML = `
    <h2 onclick="window.location.href ='profilepage.html?username=${userobj.Username}';">
    profile</h2>`
}

insertArticle(articles)

const forms = document.forms

const searchbox = forms.namedItem("searchform")

searchbox.addEventListener('submit', async(e) => {
    e.preventDefault()
    if (e.target.elements["search-input"].value === "") {
        return
    }

    const findedArticles = await fetchArticle(e.target.elements["search-input"].value)

    let finalArticles = []

    if (findedArticles["finded article"].ID !== 0) {
        finalArticles.push(findedArticles["finded article"])
    }
    
    document.getElementById("articles").innerHTML = `<h2>related articles</h2>`

    if (findedArticles["related articles"].length === 0 ) {
        document.getElementById("articles").innerHTML += 
        `<div class="article">
           <h3>no article founded by this subject</h3>
        </div>`
        return
    }

    finalArticles = [].concat(finalArticles ,findedArticles["related articles"])

    finalArticles = finalArticles.slice(0,3)

    finalArticles.forEach((element) => {
        console.log(element)
        document.getElementById("articles").innerHTML += 
        `<div class="article">
            <br>
            <h3 class="article-title">${element.Title}</h3>
            <p class="category">${element.Category}</p>
            <p class="releasedate">Date of release ${element.CreatedAt}</p>
            <p class="auther">auther 
            <a href="/front/profilepage.html?username=${element.Auther}">${element.Auther}</a>
            </p>
            <p class="desc">${element.Body}</p>
        </div>`
    })

})

