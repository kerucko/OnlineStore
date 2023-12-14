document.addEventListener("DOMContentLoaded", () => {
    const urlParams = new URLSearchParams(window.location.search);
    const id = urlParams.get('id');
    console.log(id);
    fetch("http://127.0.0.1:8080/product?id=" + id) // отправка GET-запроса на сервер
        .then(response => response.json()) // получение текстового ответа
        .then(data => {
            // JSON {name, price, description, shop}
            const name = document.getElementById("name");
            name.innerText = data.title;

            const price = document.getElementById("price");
            price.innerText = data.price + " ₽";

            const description = document.getElementById("description");
            description.innerText = data.description;

            const img = document.createElement('img');
            img.classList.add("order__card-sections--1img")
            img.src = data.photo_path
            img.alt = ""
            const container = document.getElementById("photo")
            container.append(img)

            const shop = document.getElementById("shop");
            shop.innerText = data.shop;
        })
        .catch(error => {
            console.error("Error:", error);
        });
});

function sendData() {
    const urlParams = new URLSearchParams(window.location.search);
    const id = urlParams.get('id');

    const sellerID = localStorage.getItem("id");
    console.log(sellerID)
    var data = {customer_id:sellerID, product_id:id};
    var url = 'http://localhost:8080/seller/cart'

    fetch(url, {
        mode: 'no-cors',
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    })
        .catch((error) => {
            console.error('Ошибка:', error);
        });

}