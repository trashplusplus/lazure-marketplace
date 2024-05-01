document.getElementById("add-product").addEventListener('click', function() {
    document.getElementById("overlay").style.display = 'block';
    document.getElementById("popup").style.display = 'block';
});
function toggleDropdown() {
    const dropdown = document.querySelector('.custom-select');
    dropdown.classList.toggle('open');
}

document.querySelectorAll('.custom-option').forEach(option => {
    option.addEventListener('click', function(event) {
        event.stopPropagation();

        let selectedValue = this.getAttribute('data-value');
        let selectedText = this.textContent;
        let selectTrigger = document.querySelector('.custom-select__trigger span');
        let hiddenInput = document.querySelector('input[name="type"]');

        selectTrigger.textContent = selectedText;
        hiddenInput.value = selectedValue;

        document.querySelector('.custom-select').classList.remove('open');
    });
});


window.onclick = function(e) {
    if (!e.target.matches('.custom-select, .custom-select *')) {
        document.querySelector('.custom-select').classList.remove('open');
    }
}


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
    try {
        document.getElementById('walletField').value = walletManager.getWalletString();
    } catch (error) {
        createToast("warning", "Please, connect your wallet first.");
        return;
    }

    let type = document.querySelector('input[name="type"]');
    if (!type.value) {
        createToast("warning", "Please, select product type.");
        return;
    }

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
