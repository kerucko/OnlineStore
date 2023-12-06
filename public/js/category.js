document.addEventListener("DOMContentLoaded", () => {
    const urlParams = new URLSearchParams(window.location.search);
    const category = urlParams.get('category');
    fetch("http://127.0.0.1:8080/show_all?category=" + category) // отправка GET-запроса на сервер
        .then(response => response.json()) // получение текстового ответа
        .then(data => {
            console.log(data)
            // JSON {name, [{name, price, description, shop}]}
            const category_name = document.getElementById("name");
            category_name.innerText = data.name;

            const main = document.getElementById("main");

            for (let j = 0; j < data.products.length; j += 3) {
                const row = document.createElement('div')
                row.classList.add('one__of__categories-cards')
                
                let k = Math.min(j + 3, data.products.length)
                for (let i = j; i < k; i++) {
                    const card = document.createElement('a')
                    card.setAttribute('href', 'product.html?id=' + data.products[i].id);

                    const inner_card = document.createElement('div')
                    inner_card.classList.add('one__of__categories-card')

                    const img = document.createElement('img')
                    img.classList.add('one__of__categories-img')
                    img.setAttribute('src', 'images/рубашка.png')
                    img.setAttribute('alt', 'card' + i)

                    const name = document.createElement('div')
                    name.classList.add('one__of__categories-name')
                    name.innerText = data.products[i].title

                    const price = document.createElement('div')
                    price.classList.add('one__of__categories-price')
                    price.innerText = data.products[i].price + " ₽"
 
                    inner_card.appendChild(img)
                    inner_card.appendChild(name)
                    inner_card.appendChild(price)
                    card.appendChild(inner_card)

                    row.appendChild(card)
                }

                main.appendChild(row)
            }
        })
        .catch(error => {
            console.error("Error:", error);
        });
});