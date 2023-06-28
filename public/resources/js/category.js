async function fetchArticle(searchvalue = null) {
    let response = await fetch(`/api/category/${category}/`, {method : "GET"} )
    const data = await response.json()
    data.message.forEach(element => {
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
const user = getCookie("user")

if (user !== "" ) {
    let userobj = JSON.parse(user)
    document.getElementById("user-info").innerHTML = `
    <h2 onclick="window.location.href ='profilepage.html?username=${userobj.Username}';">
    profile</h2>`
}

const url = new URL(window.location.href);
const searchParams = url.searchParams;
const category = searchParams.get('category');
document.getElementById("maintitle").innerText = category + " articles"

fetchArticle()