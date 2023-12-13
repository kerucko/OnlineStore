document.addEventListener("DOMContentLoaded", () => {
    const id = localStorage.getItem("id");
    console.log("UserID", id);
    fetch("http://127.0.0.1:8080/cart?id=" + id) // отправка GET-запроса на сервер
        .then(response => response.json()) // получение текстового ответа
        .then(data => {
            // JSON {[id, title, price, shop, photo_path]}
            const cards = document.getElementById("cards");
            for (let i = 0; i < data.length; i++) {
                cards.innerHTML += `
                <div class="basket__btns">
                    <div class="basket__cards">
                    <a href="product.html?id=${data[i].id}">
                        <div class="basket__order-card">
                            <img
                                class="basket__order-card--img"
                                src="${data[i].photo_path}"
                                alt="card${i}"
                            />
                            <div class="basket__img-name-price">
                                <div class="basket__card-name">${data[i].title}</div>
                                <div class="basket__card-price">${data[i].price} ₽</div>
                                <div class="basket__card-shop">
                                    <img class="basket__card-shop--img" src="images/магазин.svg" alt="">
                                    <div class="basket__card-shop--txt">${data[i].shop}</div>
                                </div>
                            </div>
                        </div>
                    </div>
                    </a>
                    <input type="checkbox" id="checkbox" />
                    <label for="checkbox"></label>
                </div>
                `
            }
        })
        .catch(error => {
            console.error("Error:", error);
        });
});