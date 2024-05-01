document.getElementById("add-product").addEventListener('click', function() {
    document.getElementById("overlay").style.display = 'block';
    document.getElementById("popup").style.display = 'block';
});

document.addEventListener('DOMContentLoaded', function() {
    fetch('https://productsapi-954ed826b909.herokuapp.com/category')
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
