const SOLANA_NET = 'devnet';

async function connectWallet() {
    if (window.solana && window.solana.isPhantom) {
        try {
            return await window.solana.connect({onlyIfTrusted: false});
        } catch (error) {
            console.error("Failed to connect the Phantom wallet:", error);
            throw new Error("Wallet connection failed");
        }
    } else {
        console.log("Phantom wallet not found!");
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
        const solBalance = balance / solanaWeb3.LAMPORTS_PER_SOL;
        console.log(`Balance: ${solBalance} SOL`);
        return solBalance;
    } catch (error) {
        console.error('Error getting balance:', error);
        throw new Error("Failed to get account balance");
    }
}
/*
(async () => {
    try {
        const wallet = await connectWallet();

        const balance = await getAccountBalance(wallet.publicKey);
        console.log(`The wallet balance is ${balance} SOL`);
    } catch (error) {
        console.error(error);
    }
})();*/

async function showWalletInfo() {
    const wallet = await connectWallet();

    let balance = await getAccountBalance(wallet.publicKey);
    console.log(`The wallet balance is ${balance} SOL`);

    console.log(`Wallet address is ${wallet.publicKey}`)
}
