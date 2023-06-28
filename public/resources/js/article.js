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

async function fetchArticle(articleid) {

    response = await fetch(`/api/article/${articleid}`,{ method :"GET"}) 

    if ((response.status === 400) || (response.status === 416) ){
        return {status:null, msg:"invalid or incorrect id"}
    }

    const data = await response.json()

    return {status:true, msg:data.article}

}


const url = new URL(window.location.href);

const searchParams = url.searchParams;
const articleid = searchParams.get('articleid');

const user = getCookie("user")

if (user !== "" ) {
    let userobj = JSON.parse(user)
    document.getElementById("user-info").innerHTML = `
    <h2 onclick="window.location.href ='profilepage.html?username=${userobj.Username}';">
    profile</h2>`
}

async function main() {
    if (articleid === "") {
        window.location.assign("/front/")
    }
    const result = await fetchArticle(articleid)

    if (!result.status) {
        document.getElementById("article").innerHTML += `
                <h2>invalid article id</h2>
                <p>probably meanwhile this specific article deleted</p>
            `
        return
    }

    document.getElementById("article").innerHTML = `
        <h2 class="article-title">${result.msg.Title}</h2>
        <p class="category">${result.msg.Category}</p>
        <p class="releasedate">Date of release ${result.msg.CreatedAt}</p>
        <p class="releasedate">Date of update ${result.msg.UpdatedAt}</p>
        <p class="auther">auther 
        <a href="/front/profilepage.html?username=${result.msg.Auther}">${result.msg.Auther}</a>
        </p>
        <p class="desc">${result.msg.Body}</p>
    `
    console.log(result)
}
main()



