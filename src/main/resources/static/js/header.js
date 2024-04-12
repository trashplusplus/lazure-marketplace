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

document.getElementById('close-info').addEventListener('click', function() {
    document.getElementById('wallet-info').classList.remove('open-wallet-info');
});


