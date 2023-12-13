document.addEventListener("DOMContentLoaded", () => {
    const urlParams = new URLSearchParams(window.location.search);
    const sellerID = urlParams.get('id');
    fetch("http://127.0.0.1:8080/seller/store?id=" + sellerID) // отправка GET-запроса на сервер
        .then(response => response.json()) // получение текстового ответа
        .then(data => {
            console.log(data)

            const main = document.getElementById("main");

            for (let j = 0; j < data.length; j++) {
                const row = document.createElement('div')
                row.classList.add('storage__products-card')
                const img = document.createElement('img')
                img.classList.add('seller__products-card--img')
                img.src = 'images/магазин.svg'
                img.alt = "img"
                const name = document.createElement('div')
                name.classList.add('storage__card-name')
                name.innerText = data[j].address
                row.append(img)
                row.append(name)
                main.appendChild(row)
            }

        })
        .catch(error => {
            console.error("Error:", error);
        });
});