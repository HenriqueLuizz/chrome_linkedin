function extractProfilePosts() {
    let posts = [];

    document.querySelectorAll("div.feed-shared-update-v2").forEach(post => {
        let postData = {};

        // Nome do autor
        let nameElement = post.querySelector("span.update-components-actor__single-line-truncate");
        postData.name = nameElement ? nameElement.textContent.trim() : "Name not found";

        // Link do post
        let linkElement = post.querySelector(".update-components-actor--with-control-menu");
        postData.link = linkElement ? linkElement.querySelector("a")?.href : "Link not found";
        
        // Date
        let dateElement = post.querySelector("span.update-components-actor__sub-description");
        console.log(dateElement.innerText.split("\n")[2])
        postData.date = dateElement ? dateElement.innerText.split("\n")[2] : "Date not found";
        
        // Descrição
        let descriptionElement = post.querySelector("span.update-components-actor__description");
        postData.description = descriptionElement ? descriptionElement.innerText.split('\n')[0] : "Description not found";

        // Conteúdo do post
        let textElement = post.querySelector("span.break-words");
        postData.content = textElement ? textElement.innerText.trim() : "Text not found";

        // Reações
        let reactionsElement = post.querySelector("span.social-details-social-counts__reactions-count");
        postData.reactions = reactionsElement ? reactionsElement.textContent.trim() : "0";

        // Compartilhamentos
        let repostsElement = post.querySelectorAll("button.social-details-social-counts__btn")[1];
        postData.reposts = repostsElement ? repostsElement.textContent.trim().split(" ")[0] : "0";

        posts.push(postData);
    });

    return posts;
}

// Verifica se a URL é de um perfil
function isProfilePage() {
    return window.location.pathname.startsWith("/in/");
}

// Lida com mensagens da extensão
chrome.runtime.onMessage.addListener((request, sender, sendResponse) => {
    if (request.action === "capturePosts") {
        if (isProfilePage()) {
            let posts = extractProfilePosts();
            sendResponse({ posts: posts });
        } else {
            sendResponse({ error: "Acesse um perfil do LinkedIn para capturar posts." });
        }
    }
});
