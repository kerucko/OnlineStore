document.addEventListener("DOMContentLoaded", () => {
    const id = localStorage.getItem("id");
    console.log(id);
    fetch("http://127.0.0.1:8080/all_orders?id=" + id) // отправка GET-запроса на сервер
        .then(response => response.json()) // получение текстового ответа
        .then(data => {
            // JSON {[id, status]}
           const main = document.getElementById("main");
           for (let i = 0; i < data.length; i++) {
               main.innerHTML += `
               <a href="order.html?id=${data[i].id}">
                </div>
                <div class="all__orders">
                    <div class="order__name">Заказ ${data[i].id}</div>
                    <div class="order__status">Статус заказа: ${data[i].status}</div>
                </div>
                </a>
               `
           }
        })
        .catch(error => {
            console.error("Error:", error);
        });
});