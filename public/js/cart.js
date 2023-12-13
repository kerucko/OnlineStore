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
                        <div class="basket__order-card">
                        <a href="product.html">
                            <img
                            class="basket__order-card--img"
                            src="images/рубашка.png"
                            alt="card1"
                            />
                        </a>
                        <div class="basket__img-name-price">
                            <div class="basket__card-name">
                            Рубашка женская oversize
                            </div>
                            <div class="basket__card-price">999 ₽</div>
                            <div class="basket__card-shop">
                            <img
                                class="basket__card-shop--img"
                                src="images/магазин.svg"
                                alt=""
                            />
                            <div class="basket__card-shop--txt">
                                Название магазина
                            </div>
                            </div>
                        </div>
                        <div class="cart__quantity">
                            <label for="quantity">Количество:</label>
                            <div class="cart_btns">
                            <button
                                class="btn_q"
                                type="button"
                                onclick="this.nextElementSibling.stepDown()"
                            >
                                -
                            </button>
                            <input
                                type="number"
                                min="1"
                                max="100"
                                value="1"
                                readonly
                                class="raz"
                            />
                            <button
                            class="btn_q"
                                type="button"
                                onclick="this.previousElementSibling.stepUp()"
                            >
                                +
                            </button>
                            </div>
                        </div>
                        </div>
                    </div>
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