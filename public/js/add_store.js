function sendData() {
    const urlParams = new URLSearchParams(window.location.search);
    const sellerID = urlParams.get('id');
    var textValue = document.getElementById('text').value;
    console.log(sellerID)
    var data = {id:0, address: textValue };
    var url = 'http://localhost:8080/seller/store?id=' + sellerID;

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
    location.href="seller_storage.html?id=" + sellerID
}