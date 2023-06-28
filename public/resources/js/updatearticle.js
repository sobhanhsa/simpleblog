async function getCurrentArticle(articleid, username) {

    belongsto = false

    let response
    
    response = await fetch(`/api/article/${articleid}`,{ method :"GET"}) 

    if ((response.status === 400) || (response.status === 416) ){
        return {status:null, msg:"invalid id"}
    }

    const data = await response.json()

    if (data.article.Auther !== username ) {
        return {status : false, msg:"you dont have acces to this article"}
    }

    return {status : true, msg:data.article}

}

async function updateArticle(articleid,{title , category, body, hashtag}) {

    const response = await fetch(`/api/article/${articleid}`, {
        method: "PUT", // *GET, POST, PUT, DELETE, etc.
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

    return {status : true, msg : data.message.article}
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



async function main() {
    
    let form = document.querySelector('#update-form')

    let user = getCookie("user");

    if (user === "") {
        document.getElementsByClassName("update-form")[0].innerHTML =  `
        <h2>you are not logged in, click 
        <a href="/loginpage.html">here</a> to login</h2>
        `
        return
    }
    
    const userobj = JSON.parse(user)
    
    const url = new URL(window.location.href);
    
    const searchParams = url.searchParams;
    const articleid = searchParams.get('id');
    
    const access = await getCurrentArticle(articleid, userobj.Username)

    if (access.status === null) {
        document.getElementsByClassName("update-form")[0].innerHTML =  `
        <h2>invalid article id.
        <a href="/">home</a></h2>
        `
        return
    }

    if (access.status === false) {
        document.getElementsByClassName("update-form")[0].innerHTML =  `
        <h2>you dont have access to this article.
        <a href="/">home</a></h2>
        `
        return
    }

    //standardized elements
    document.getElementsByClassName("content")[0]. innerHTML = `
    <div class="input-field">
        <input type="text" placeholder="Title( ${access.msg.Title} )" id="title-field">
        <p id="errtitle"></p>
    </div>
    <div class="category-field">
        <select name="caregory" id="category" form="update-form">
            <option value="sport">Sport</option>
            <option value="business">Business</option>
            <option value="science">Science</option>
            <option value="lifestyle" selected>LifeStyle</option>
        </select>
    </div>
    <div class="input-field">
        <input type="text" placeholder="body(${access.msg.Body.substring(0,20)}...)" id="body-field" >
        <p id="errbody"></p>
    </div>

    <div class="input-field">
        <input type="text" placeholder="Hash tag(${access.msg.HashTag})" id="hashtag-field">
        <p id="errhashtag"></p>
    </div>
    `

    const $select = document.querySelector('#category');
    $select.value = access.msg.Category

    form.addEventListener('submit', async(e) => {

        e.preventDefault();

        if (e.submitter.innerText === "CANCEL") {
            window.location.assign("/")
            return
        } else if (e.submitter.innerText === "UPDATE ARTICLE") {

            const title = e.target.elements["title-field"].value

            const category = e.target.elements["category"].value
    
            const body = e.target.elements["body-field"].value
    
            const hashtag = e.target.elements["hashtag-field"].value

            if ((title === "") || (body === "") ) {
                alert("please input required fields")
                return
            }
    
            const result = await updateArticle(articleid,{title, category, body, hashtag})

            if (result.status === false) {
                document.getElementById("errtitle").innerText = "taken Title"
                return
            }

            document.getElementsByClassName("update-form")[0].innerHTML =  `
                <h2>you're you're article successfully updated.</h2>
            `

            setTimeout(function(){
                console.log("after 1 second")
                window.location.assign(`/profilepage.html?username=${userobj.Username}`)
            }, 1500);

        }

    })

}
main()



