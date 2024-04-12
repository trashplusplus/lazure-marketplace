const SOLANA_NET = 'devnet';
const LOCALSTORAGE_USER_REJECTED_ID = 'userRejectedWalletConnection';

(async () => {
    try {
        if (!getUserRejectedRequest()) {
            const wallet = await connectWallet();
            document.getElementById("balance").innerText = await getAccountBalance(wallet.publicKey) + " SOL";
        }
    } catch (error) {
        setUserRejected(true);
    }
})();


async function connectWallet() {
    if (window.solana && window.solana.isPhantom) {
        try {
            let wallet = await window.solana.connect({onlyIfTrusted: false});
            document.getElementById("balance").innerText = await getAccountBalance(wallet.publicKey) + " SOL";
            setUserRejected(false);
            return wallet;
        } catch (error) {
            createToast("warning", `Failed to connect the Phantom wallet: ${error}`);
            setUserRejected(true);
            throw new Error("Wallet connection failed");
        }
    } else {
        createToast("info", "You should install Phantom wallet first.");
        throw new Error("Phantom wallet not found");
    }
}

function getSolanaConnection() {
    return new solanaWeb3.Connection(solanaWeb3.clusterApiUrl(SOLANA_NET), 'confirmed');
}

async function getAccountBalance(publicKey) {
    try {
        const connection = getSolanaConnection();
        const balance = await connection.getBalance(publicKey);
        return balance / solanaWeb3.LAMPORTS_PER_SOL;
    } catch (error) {
        createToast("error", `Error getting balance: ${error}`);
        throw new Error("Failed to get account balance");
    }
}

async function showWalletInfo() {
    const wallet = await connectWallet();

    let balance = await getAccountBalance(wallet.publicKey);
    let shortWalletAddress = shortenWalletAddress(wallet.publicKey.toString());
    document.getElementById("wallet-info").classList.add("open-wallet-info");
    document.getElementById("wallet-short-info").innerText = `${shortWalletAddress} ${balance} SOL`;
}

function setUserRejected(isRejected) {
    localStorage.setItem(LOCALSTORAGE_USER_REJECTED_ID, isRejected.toString());
}

function getUserRejectedRequest() {
    return localStorage.getItem(LOCALSTORAGE_USER_REJECTED_ID) === "true";
}

function shortenWalletAddress(fullAddress) {
    if (fullAddress.length > 8) {
        return `${fullAddress.slice(0, 5)}...${fullAddress.slice(-3)}`;
    }
    return fullAddress;
}

async function disconnectWallet() {
    if (window.solana && window.solana.isPhantom) {
        try {
            await window.solana.disconnect();
            createToast("info", "Wallet was successfully disconnected!");
            document.getElementById("balance").innerText = "Connect Wallet";
            document.getElementById('wallet-info').classList.remove('open-wallet-info');
        } catch (error) {
            createToast("error", "Error disconnecting wallet:", error)
        }
    }
}