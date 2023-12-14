document.addEventListener("DOMContentLoaded", () => {
    const urlParams = new URLSearchParams(window.location.search);
    const id = urlParams.get('id');
    console.log("OrderID", id);
    fetch("http://127.0.0.1:8080/order?id=" + id) // отправка GET-запроса на сервер
        .then(response => response.json()) // получение текстового ответа
        .then(data => {
            console.log(data)
            // JSON {id, status, [id, title, price, shop, photo_path]}
            const cards = document.getElementById("cards");
            for (let i = 0; i < data.products.length; i++) {
                console.log(i);
                cards.innerHTML += `
                <a class="orders" href="product.html?id=${data.products[i].id}">
                <div class="basket__order-card">
                  <img
                    class="basket__order-card--img"
                    src="${data.products[i].photo_path}"
                    alt="card${i}"
                  />
                  <div class="basket__img-name-price">
                  <div class="basket__card-name">${data.products[i].title}</div>
                  <div class="basket__card-price">
                    ${data.products[i].price} ₽
                  </div>
                  <div class="basket__card-shop">
                      <img class="basket__card-shop--img" src="images/магазин.svg" alt="">
                      <div class="basket__card-shop--txt">${data.products[i].shop}</div>
                  </div>
                  <div class="basket__card-shop--quantity">Количество: ${data.products[i].amount}</div>   
              </div>
                </div>
              </a>
                `
            }
        })
        .catch(error => {
            console.error("Error:", error);
        });
});