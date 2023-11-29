document.addEventListener("DOMContentLoaded", () => {
    const urlParams = new URLSearchParams(window.location.search);
    const id = urlParams.get('id');
    console.log(id);
    fetch("http://127.0.0.1:8000/product/" + id) // отправка GET-запроса на сервер
        .then(response => response.json()) // получение текстового ответа
        .then(data => {
            // JSON {name, price, description, shop}
            const name = document.getElementById("name");
            name.innerText = data.name;

            const price = document.getElementById("price");
            price.innerText = data.price + " ₽";

            const description = document.getElementById("description");
            description.innerText = data.description;

            const shop = document.getElementById("shop");
            shop.innerText = data.shop;
        })
        .catch(error => {
            console.error("Error:", error);
        });
});