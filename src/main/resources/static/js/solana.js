const SOLANA_NET = 'devnet';
const LOCALSTORAGE_USER_REJECTED_ID = 'userRejectedWalletConnection';

let subscribers = [];

class WalletManager {
    constructor() {
        this.wallet = null;
        this.connection = null;
        this.init();
    }

    async init() {
        try {
            if (!this.getUserRejectedRequest()) {
                this.wallet = await this.connectWallet();
                this.connection = this.getSolanaConnection();
                await this.updateDisplayedUserInfo();
            }
        } catch (error) {
            console.error("Initialization failed:", error);
        }
    }

    onWalletReady(callback) {
        if (this.wallet) {
            callback();
        } else {
            subscribers.push(callback);
        }
    }

    notifySubscribers() {
        if (!this.wallet) {
            createToast("warning", "Wallet not ready yet.");
            return;
        }
        subscribers.forEach(callback => callback(this.getWalletString()));
        subscribers = [];
    }

    getSolanaConnection() {
        return new solanaWeb3.Connection(solanaWeb3.clusterApiUrl(SOLANA_NET), 'confirmed');
    }

    sendLoginRequest() {
        fetch('/api/users/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ walletId: this.getWalletString() })
        })
            .then(response => {
                if (!response.ok) {
                    throw new Error('Network response was not ok: ' + response.statusText);
                }
            })
            .catch(error => {
                throw new Error("Error while getting login data.");
            });
    }

    async connectWallet() {
        if (window.solana && window.solana.isPhantom) {
            try {
                this.wallet = await window.solana.connect({onlyIfTrusted: false});
                this.setUserRejected(false);
                await this.updateDisplayedUserInfo();
                this.sendLoginRequest();

                this.notifySubscribers();

                return this.wallet;
            } catch (error) {
                this.handleWalletConnectionError(error);
                throw error;
            }
        } else {
            createToast("info", "You should install Phantom wallet first.");
            throw new Error("Phantom wallet not found");
        }
    }

    async updateDisplayedUserInfo() {
        let wallet = this.wallet.publicKey;
        let profileURL = document.getElementById("profile");
        const balance = await this.getAccountBalance(wallet);

        document.getElementById("balance").innerText = `${balance} SOL`;

        profileURL.innerText = this.shortenWalletAddress(wallet.toString());
        profileURL.href = "/profile/" + wallet.toString();
    }

    async getAccountBalance(publicKey) {
        if (!this.connection) {
            this.connection = this.getSolanaConnection();
        }
        try {
            const balance = await this.connection.getBalance(publicKey);
            return balance / solanaWeb3.LAMPORTS_PER_SOL;
        } catch (error) {
            createToast("error", `Error getting balance: ${error}`);
            throw new Error("Failed to get account balance");
        }
    }

    getWalletString() {
        return this.wallet.publicKey.toString();
    }

    async showWalletInfo() {
        if (!this.wallet) {
            await this.connectWallet();
        }

        document.getElementById("wallet-info").classList.add("open-wallet-info");
        document.getElementById("wallet-short-info").innerText = "Loading...";

        try {
            const balance = await this.getAccountBalance(this.wallet.publicKey);
            const shortWalletAddress = this.shortenWalletAddress(this.getWalletString());
            document.getElementById("wallet-short-info").innerText = `${shortWalletAddress} ${balance} SOL`;
        } catch (error) {
            document.getElementById("wallet-short-info").innerText = "Failed to load data";
            console.error("Failed to load wallet info:", error);
        }
    }


    shortenWalletAddress(fullAddress) {
        return fullAddress.length > 8 ? `${fullAddress.slice(0, 5)}...${fullAddress.slice(-3)}` : fullAddress;
    }

    async disconnectWallet() {
        if (window.solana && window.solana.isPhantom) {
            try {
                await window.solana.disconnect();
                createToast("info", "Wallet was successfully disconnected!");
                document.getElementById("balance").innerText = "Connect Wallet";
                document.getElementById('wallet-info').classList.remove("open-wallet-info");
                let profileURL = document.getElementById("profile");
                profileURL.innerText = "";
                profileURL.href = "";
                this.wallet = null;
                this.connection = null;
                this.setUserRejected(true);
            } catch (error) {
                createToast("error", `Error disconnecting wallet: ${error}`);
            }
        }
    }

    setUserRejected(isRejected) {
        localStorage.setItem(LOCALSTORAGE_USER_REJECTED_ID, isRejected.toString());
    }

    getUserRejectedRequest() {
        return localStorage.getItem(LOCALSTORAGE_USER_REJECTED_ID) === "true";
    }

    handleWalletConnectionError(error) {
        createToast("warning", `Failed to connect the Phantom wallet: ${error}`);
        this.setUserRejected(true);
    }
}

const walletManager = new WalletManager();
document.getElementById('profile-balance').addEventListener('click', function() {
    walletManager.showWalletInfo();
});

document.getElementById('disconnect-wallet').addEventListener('click', function() {
    walletManager.disconnectWallet();
});