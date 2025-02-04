document.getElementById("captureBtn").addEventListener("click", () => {
    chrome.tabs.query({ active: true, currentWindow: true }, (tabs) => {
        chrome.scripting.executeScript({
            target: { tabId: tabs[0].id },
            function: extractPostsFromProfile
        }, (results) => {
            if (results && results[0] && results[0].result.length > 0) {
                let posts = results[0].result;
                document.getElementById("status").innerText = "Posts captured!";
                
                // Stores the posts for subsequent export
                window.collectedPosts = posts;
                
                document.getElementById("downloadJson").disabled = false;
            } else {
                document.getElementById("status").innerText = "No posts found or incorrect page!";
            }
        });
    });
});

document.getElementById("downloadJson").addEventListener("click", () => {
    if (window.collectedPosts) {
        saveAsJSON(window.collectedPosts);
    }
});

function getTimestamp() {
    const now = new Date();
    const year = now.getFullYear();
    const month = String(now.getMonth() + 1).padStart(2, '0');
    const day = String(now.getDate()).padStart(2, '0');

    return `${year}-${month}-${day}`;
}

function extractPostsFromProfile() {
    const now = new Date();
    const year = now.getFullYear();
    const month = String(now.getMonth() + 1).padStart(2, '0');
    const day = String(now.getDate()).padStart(2, '0');
    const hours = String(now.getHours()).padStart(2, '0');
    const minutes = String(now.getMinutes()).padStart(2, '0');
    const seconds = String(now.getSeconds()).padStart(2, '0');

    if (!window.location.pathname.startsWith("/in/")) {
        return [];
    }

    let posts = [];
    document.querySelectorAll("div.feed-shared-update-v2").forEach(post => {
        let postData = {};

        // Actor Name
        let nameElement = post.querySelector("span.update-components-actor__single-line-truncate");
        postData.name = nameElement ? nameElement.textContent.trim() : "Name not found";

        // Actor Link
        let linkElement = post.querySelector(".update-components-actor--with-control-menu");
        postData.link = linkElement ? linkElement.querySelector("a")?.href : "Link not found";
        
        // Post Date
        let dateElement = post.querySelector("span.update-components-actor__sub-description");
        postData.date = dateElement ? dateElement.innerText.split("\n")[2] : "Date not found";

        // Extraction Date
        postData.timestamp = `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`;
        
        // Description
        let descriptionElement = post.querySelector("span.update-components-actor__description");
        postData.description = descriptionElement ? descriptionElement.innerText.split('\n')[0] : "Description not found";

        // Post Content
        let textElement = post.querySelector("span.break-words");
        postData.content = textElement ? textElement.innerText.trim() : "Text not found";

        // Reactions
        let reactionsElement = post.querySelector("span.social-details-social-counts__reactions-count");
        postData.reactions = reactionsElement ? reactionsElement.textContent.trim() : "";

        // Reposts
        let repostsElement = post.querySelector("button.social-details-social-counts__btn span");
        postData.reposts = repostsElement ? repostsElement.textContent.trim().split(" ")[0] : "";

        posts.push(postData);
    });

    return posts;
}

function saveAsJSON(posts) {
    const jsonData = JSON.stringify(posts, null, 2);
    const blob = new Blob([jsonData], { type: "application/json" });
    const url = URL.createObjectURL(blob);

    chrome.downloads.download({
        url: url,
        filename: `linkedin_posts_${getTimestamp()}.json`,
        saveAs: true
    }, () => URL.revokeObjectURL(url));
}
