function sendData() {
    const sellerID = localStorage.getItem("id");
    var title = document.getElementById('name').value;
    var category = document.getElementById('category').value
    var price = document.getElementById('price').value
    var desc = document.getElementById('description').value
    var address = document.getElementById('address').value
    var amount = document.getElementById('amount').value
    console.log(sellerID)
    var data = {
        title: title,
        category: category,
        price: parseInt(price),
        description: desc,
        store_address: address,
        amount: parseInt(amount)
    };
    var url = 'http://localhost:8080/seller/product?id=' + sellerID;

    fetch(url, {
        mode: 'no-cors',
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    })
        .catch((error) => {
            console.error('Ошибка:', error);
        });
    location.href="seller_products.html"
}