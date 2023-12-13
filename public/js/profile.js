document.addEventListener("DOMContentLoaded", () => {
    const id = localStorage.getItem("id");
    console.log("UserID", id);
    fetch("http://127.0.0.1:8080/profile?id=" + id) // отправка GET-запроса на сервер
        .then(response => response.json()) // получение текстового ответа
        .then(data => {
            // JSON {id, name, phone, email, address}
            const name = document.getElementById("name");
            name.value = data.name

            const phone = document.getElementById("tel");
            phone.value = data.phone

            const email = document.getElementById("email");
            email.value = data.email

            const address = document.getElementById("address");
            address.value = data.address
        })
        .catch(error => {
            console.error("Error:", error);
        });
});