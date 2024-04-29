document.getElementById("add-product").addEventListener('click', function() {
    document.getElementById("overlay").style.display = 'block';
    document.getElementById("popup").style.display = 'block';
});

document.getElementById("overlay").addEventListener('click', function() {
    document.getElementById("overlay").style.display = 'none';
    document.getElementById("popup").style.display = 'none';
});

document.getElementById("close-popup").addEventListener('click', function() {
    document.getElementById("overlay").style.display = 'none';
    document.getElementById("popup").style.display = 'none';
});

document.querySelector('form').addEventListener('submit', function(event) {
    event.preventDefault();
    document.getElementById('walletField').value = walletManager.getWalletString();

    let form = event.target;
    let data = new FormData(form);

    fetch(form.action, {
        method: 'POST',
        body: data
    })
        .then(response => response.text())
        .then(text => {
            createToast("success", text);
        })
        .catch(error => {
            createToast("warning", error);
        });
});
