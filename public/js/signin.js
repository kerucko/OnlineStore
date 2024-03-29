function signin() {
    const email = document.getElementById("email").value;
    const password = document.getElementById("password").value;
    console.log(email)
    console.log(password)

    const isSeller = document.getElementById("checkbox").checked;

    fetch("http://127.0.0.1:8080/signin?email=" + email + "&password=" + password)
    .then(response => response.json())
    .then(data => {
        console.log(data)
        localStorage.setItem("id", data.id)
        if (isSeller) {
            location.href = "seller.html"
        } else {
            location.href = "main.html"
        }
    })
    .catch(error => {
        console.error("Error:", error);
    });
}
