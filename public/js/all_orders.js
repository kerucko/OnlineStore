document.addEventListener("DOMContentLoaded", () => {
    const urlParams = new URLSearchParams(window.location.search);
    const id = urlParams.get('id');
    console.log(id);
    fetch("http://127.0.0.1:8080/product?id=" + id) // отправка GET-запроса на сервер
        .then(response => response.json()) // получение текстового ответа
        .then(data => {
            // JSON {name, price, description, shop}
           
        })
        .catch(error => {
            console.error("Error:", error);
        });
});