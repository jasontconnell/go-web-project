
window.addEventListener("DOMContentLoaded", async () => {
    var test = document.querySelector("#test");
    var resp = await fetch("/api/test");
    var json = await resp.json();
    test.innerHTML = json.Message;
})