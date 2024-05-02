document.getElementById("add-product").addEventListener('click', function() {
    document.getElementById("overlay").style.display = 'block';
    document.getElementById("popup").style.display = 'block';
});

document.addEventListener('DOMContentLoaded', function() {
    fetch('api/products/category')
        .then(response => {
            if (!response.ok) {
                throw new Error('Network response was not ok: ' + response.statusText);
            }
            return response.json();
        })
        .then(data => {
            const container = document.querySelector('.custom-options');
            container.innerHTML = '';

            data.forEach(category => {
                const option = document.createElement('span');
                option.className = 'custom-option';
                option.setAttribute('data-value', category.category_id);
                option.textContent = category.name;

                const infoIcon = document.createElement('img');
                infoIcon.src = 'img/info.png';
                infoIcon.alt = 'Info';
                infoIcon.className = 'info-icon';
                infoIcon.title = category.description;

                option.appendChild(infoIcon);

                option.addEventListener('click', function() {
                    let selectedValue = this.getAttribute('data-value');
                    let hiddenInput = document.querySelector('input[name="categoryId"]');
                    hiddenInput.value = selectedValue;
                    document.querySelectorAll('.custom-option').forEach(opt => opt.classList.remove('selected'));
                    this.classList.add('selected');
                    document.querySelector('.custom-select__trigger span').textContent = this.textContent;
                });

                container.appendChild(option);
            });
        })
        .catch(error => {
            console.error('Failed to fetch categories:', error);
        });
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
        let hiddenInput = document.querySelector('input[name="categoryId"]');

        selectTrigger.textContent = selectedText;
        hiddenInput.value = selectedValue;
        console.log(selectedValue);
        document.querySelector('.custom-select').classList.remove('open');
    });
});


window.onclick = function(e) {
    if (!e.target.matches('.custom-select, .custom-select *')) {
        document.querySelector('.custom-select').classList.remove('open');
    }
}


document.querySelector('form').addEventListener('submit', function(event) {
    event.preventDefault();

    let categoryId = document.querySelector('input[name="categoryId"]');
    if (!categoryId.value) {
        createToast("warning", "Please, select product type.");
        return;
    }

    let form = event.target;
    let data = new FormData(form);

    fetch(form.action, {
        method: 'POST',
        body: data
    })
        .then(response => {
            if (!response.ok) {
                return response.text().then(text => Promise.reject(text));
            }
            return response.text();
        })
        .then(text => {
            createToast("success", text);
            loadListings();
        })
        .catch(error => {
            createToast("warning", error);
        });

});

document.getElementById("overlay").addEventListener('click', function() {
    document.getElementById("overlay").style.display = 'none';
    document.getElementById("popup").style.display = 'none';
});

document.getElementById("close-popup").addEventListener('click', function() {
    document.getElementById("overlay").style.display = 'none';
    document.getElementById("popup").style.display = 'none';
});

document.addEventListener('DOMContentLoaded', loadListings);

function loadListings() {
    const loader = document.querySelector('.loader');
    const container = document.querySelector('.main-products-container');

    walletManager.onWalletReady((wallet) => {
        fetch('api/products/wallet/' + wallet)
            .then(response => response.json())
            .then(data => {
                loader.style.display = 'none';
                container.innerHTML = '';
                if (data.length === 0) {
                    const img = document.createElement('img');
                    img.src = "/img/broken-image.png";
                    img.alt = "no image lol";
                    const productDiv = document.createElement('div');
                    productDiv.className = 'product';
                    const infoDiv = document.createElement('div');
                    infoDiv.className = 'short-product-info';
                    const nameP = document.createElement('p');
                    nameP.textContent = "No listings found";
                    infoDiv.appendChild(nameP);
                    productDiv.appendChild(img);
                    productDiv.appendChild(infoDiv);

                    container.appendChild(productDiv);
                } else {
                    data.forEach(product => {
                        container.appendChild(createProductElement(product));
                    });
                }
            })
            .catch(error => {
                loader.style.display = 'none';
                createToast("error", "Failed to load products.");
            });
    });
}


function createProductElement(product) {
    const productDiv = document.createElement('div');
    productDiv.className = 'product';

    const imageMap = {
        1: 'code.png',
        2: 'asset.png',
        3: 'intellectual-property.png',
        4: 'coupon.png'
    };

    const img = document.createElement('img');
    img.src = `/img/${imageMap[product.category_id] || 'logo.png'}`;
    img.alt = product.name;

    const infoDiv = document.createElement('div');
    infoDiv.className = 'short-product-info';

    const nameP = document.createElement('p');
    nameP.id = 'name';
    nameP.textContent = product.name;

    const priceP = document.createElement('p');
    const priceSpan = document.createElement('span');
    priceSpan.className = 'price';
    priceSpan.textContent = product.price.toFixed(2);
    priceP.appendChild(priceSpan);
    priceP.append(' SOL');

    infoDiv.appendChild(nameP);
    infoDiv.appendChild(priceP);
    productDiv.appendChild(img);
    productDiv.appendChild(infoDiv);

    return productDiv;
}
