document.getElementById("configure-search-button").addEventListener("click", function () {
    document.getElementById('search-config').classList.add('open-search-config');
});

document.getElementById("close-search-config").addEventListener("click", function () {
    document.getElementById('search-config').classList.remove('open-search-config');
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
            const container = document.getElementById('search-form-checkboxes');
            container.innerHTML = '';

            data.forEach(category => {
                const paragraphElement = document.createElement('p');
                paragraphElement.className = 'search-config-form-item';

                const spanElement = document.createElement('span');
                spanElement.textContent = category.name;
                spanElement.title = category.description;

                const checkbox = document.createElement('input');
                checkbox.type = "checkbox";
                checkbox.className = "search-config-property";
                checkbox.name = "selectedCategory";
                checkbox.value = category.category_id;
                document.getElementById('search-form-checkboxes').appendChild(checkbox);


                paragraphElement.appendChild(checkbox);
                paragraphElement.appendChild(spanElement);


                container.appendChild(paragraphElement);
            });
        })
        .catch(error => {
            console.error('Failed to fetch categories:', error);
        });
});

document.getElementById('search-config-form').addEventListener('submit', function(event) {
    event.preventDefault();
    let form = event.target;

    let url = new URL(form.action);
    let params = new URLSearchParams(new FormData(form));

    url.search = params.toString();

    fetch(url, {
        method: 'GET',
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
