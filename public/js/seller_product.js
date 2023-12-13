document.addEventListener("DOMContentLoaded", () => {
    const sellerID = localStorage.getItem("id");
    fetch("http://127.0.0.1:8080/seller/product?id=" + sellerID) // отправка GET-запроса на сервер
        .then(response => response.json()) // получение текстового ответа
        .then(data => {
            const main = document.getElementById("main");
            for (let i = 0; i < data.length; i++) {
                main.innerHTML += `
                <div class="seller__products-card">
                    <img
                      class="seller__products-card--img"
                      src = "${data[i].photo_path}"
                      alt="product"
                    />
                    <div class="seller__img-name-price">
                      <div class="seller__name-storage">
                        <div class="seller__card-name">${data[i].title}</div>
                        <div class="seller__card-storage--txt">Склад: ${data[i].store_address}</div>
                      </div>
                    </div>
                    <div class="seller__card-quantity">
                      Количество <br> на складе:
                      <div class="seller__card-quantity--num">${data[i].amount}</div>
                    </div>
                </div>
                `
            }

        })
        .catch(error => {
            console.error("Error:", error);
        });
});