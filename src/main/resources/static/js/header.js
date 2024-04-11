document.getElementById('logo').onclick = function () {
    location.href='/'
};

document.querySelectorAll('.side-menu-element').forEach(function(element) {
    element.addEventListener('click', function() {
        let link = this.querySelector('a');
        window.location.href = link.getAttribute('href');
    });
});

document.getElementById('menu').addEventListener('click', function() {
    document.getElementById('side-menu').classList.add('open-menu');
});

document.getElementById('close-menu').addEventListener('click', function() {
    document.getElementById('side-menu').classList.remove('open-menu');
});


function closeDropdown() {
    let dropdown = document.querySelector('.language-dropdown .dropdown');
    dropdown.style.display = 'none';
}

