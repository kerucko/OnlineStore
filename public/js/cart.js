document.addEventListener("DOMContentLoaded", () => {
    const id = localStorage.getItem("id");
    console.log("UserID", id);
    fetch("http://127.0.0.1:8080/cart?id=" + id) // отправка GET-запроса на сервер
        .then(response => response.json()) // получение текстового ответа
        .then(data => {
            // JSON {[id, title, price, shop, photo_path, amount]}
            const cards = document.getElementById("cards");
            for (let i = 0; i < data.length; i++) {
                cards.innerHTML += `
                <div class="basket__btns">
              <div class="basket__cards">
                <div class="basket__order-card">
                  <a href="product.html?id=${data[i].id}">
                    <img
                      class="basket__order-card--img"
                      src="${data[i].photo_path}"
                      alt="card${i}"
                    />
                  </a>
                  <div class="basket__img-name-price">
                    <div class="basket__card-name">
                      ${data[i].title}
                    </div>
                    <div class="basket__card-price">${data[i].price} ₽</div>
                    <div class="basket__card-shop">
                      <img
                        class="basket__card-shop--img"
                        src="images/магазин.svg"
                        alt=""
                      />
                      <div class="basket__card-shop--txt">
                        ${data[i].shop}
                      </div>
                    </div>
                  </div>
                  <div class="cart__del-quant">
                    <button class="cart__delete" onclick="del(${data[i].id})">Удалить</button>

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
                          value="${data[i].amount}"
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
              </div>
            </div>
                `
            }
        })
        .catch(error => {
            console.error("Error:", error);
        });
});

function del(product_id) {
    const customer_id = localStorage.getItem("id");
    fetch(`http://127.0.0.1:8080/delete_from_cart?customer_id=${customer_id}&product_id=${product_id}`)
    .then(response => response.json()) // получение текстового ответа
        .then(data => {
            if (data == true) {
                location.reload();
            } else {
                console.error("Error:", data);
            }
        })
        .catch(error => {
            console.error("Error:", error);
        });
}

function buy() {
    const id = localStorage.getItem("id");
    fetch(`http://127.0.0.1:8080/buy?id=${id}`, {
        mode: 'no-cors',
        method: 'POST'
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        return response.json();
    })
    .then(data => {
        console.log('Item added to orders:', data);
        // Refresh the cart or handle the UI update here
    })
    .catch(error => {
        console.error('Error:', error);
    });
}