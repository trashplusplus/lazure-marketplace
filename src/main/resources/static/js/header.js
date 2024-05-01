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

document.addEventListener('click', function(event) {
    const sideMenu = document.getElementById('side-menu');
    const walletInfo = document.getElementById('wallet-info');

    if (!sideMenu.contains(event.target) && sideMenu.classList.contains('open-menu')) {
        sideMenu.classList.remove('open-menu');
    }

    if (!walletInfo.contains(event.target) && walletInfo.classList.contains('open-wallet-info')) {
        walletInfo.classList.remove('open-wallet-info');
    }
});

document.getElementById('menu').addEventListener('click', function(event) {
    document.getElementById('side-menu').classList.add('open-menu');
    event.stopPropagation();
});

document.getElementById('profile-balance').addEventListener('click', function(event) {
    document.getElementById('wallet-info').classList.add('open-wallet-info');
    event.stopPropagation();
});

