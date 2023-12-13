document.addEventListener("DOMContentLoaded", () => {
    const urlParams = new URLSearchParams(window.location.search);
    const sellerID = urlParams.get('id');
    fetch("http://127.0.0.1:8080/seller/delivery?id=" + sellerID) // отправка GET-запроса на сервер
        .then(response => response.json()) // получение текстового ответа
        .then(data => {
            const main = document.getElementById("main");
            for (let i = 0; i < data.length; i++) {
                main.innerHTML += `
                <div class="seller__products-card">
                    <img
                        class="seller__products-card--img"
                        src="${data[i].photo_path}"
                        alt="card1"
                    />
                    <div class="seller__img-name-price">
                        <div class="seller__name-storage">
                            <div class="seller__card-name">${data[i].title}</div>
                            <div class="seller__card-date">Дата поставки:
                                <div class="seller__card-date1">${data[i].data.substring(0, 10)}</div>
                            </div>
                            <div class="seller__card-delivery--txt">Склад: ${data[i].store_address}</div>
                        </div>
                    </div>
                    <div class="seller__card-quantity">
                        Количество <br> в поставке:
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

