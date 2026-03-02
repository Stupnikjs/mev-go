package main

import (
	"github.com/ethereum/go-ethereum/common"
)

// GetPools reserve

var ADDRESS_TO_SYMBOL_MAINNET = map[common.Address]string{
	// --- NATIVE / WRAPPED ---
	common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"): "WETH",
	common.HexToAddress("0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C599"): "WBTC",

	// --- STABLECOINS ---
	common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eb48"): "USDC",
	common.HexToAddress("0xdAC17F958D2ee523a2206206994597C13D831ec7"): "USDT",
	common.HexToAddress("0x6B175474E89094C44Da98b954EedeAC495271d0F"): "DAI",
	common.HexToAddress("0x1aBaEA1f7230f38fFB5A03b8395B4d3f1fd3872E"): "EURA",
	common.HexToAddress("0x6c3ea9036406852006290770bedfc107f186715f"): "PYUSD",

	// --- BLUE CHIPS / DEFI ---
	common.HexToAddress("0x514910771AF9Ca656af840dff83E8264EcF986CA"): "LINK",
	common.HexToAddress("0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984"): "UNI",
	common.HexToAddress("0x7Fc66500c84A76Ad7e9c93437bFc5Ac33E2DDaE9"): "AAVE",
	common.HexToAddress("0x5A987829255094551674823AeeD7103027f156f8"): "LDO",
	common.HexToAddress("0x9f8F72aA9304c8B593d555F12eF6589cC3A579A2"): "MKR",
	common.HexToAddress("0xc00e94cb662c3520282e6f5717214004a7f26888"): "COMP",

	// --- LAYER 2 / BRIDGE ---
	common.HexToAddress("0xB50721BCf8d664c30412Cfbc6cf7a15145234ad1"): "ARB",
	common.HexToAddress("0x4200000000000000000000000000000000000042"): "OP",
	common.HexToAddress("0x594899599f4235B2E7144256444655033d940EFd"): "USDT",
	common.HexToAddress("0xFd6254c7b032b5cE940903729172919C56805465"): "WLD",
	common.HexToAddress("0x14feE680690900BA0ccCfC76AD70Fd1b95D10e16"): "PAAL",
}

var ADDRESS_TO_SYMBOL_SEPOLIA = map[common.Address]string{
	// --- NATIVE / WRAPPED ---
	common.HexToAddress("0xfff9976782d46cc05630d1f6ebab18b2324d6b14"): "WETH", // Wrapped Ether officiel sur Sepolia (Uniswap + la plupart des apps)

	// --- STABLECOINS ---
	common.HexToAddress("0x1c7D4B196Cb0C7B01d743Fbc6116a902379C7238"): "USDC", // USDC officiel Circle sur Sepolia (le plus utilisé pour tests DeFi)
	common.HexToAddress("0xff34b3d4aee8ddcd6f9afffb6fe49bd371b8a357"): "DAI",  // DAI test (souvent mint via Aave faucet ou forks)
	// USDT : pas de déploiement "officiel" ultra-standard comme sur mainnet, mais souvent testé via Aave ou Uniswap pools → utilise un mock si besoin
	// WBTC : similaire, souvent mock via Aave faucet Sepolia (pas d'adresse "canonique" fixe comme mainnet)

	// --- BLUE CHIPS / DEFI (test versions) ---
	common.HexToAddress("0x779877A7B0D9E8603169DdbD7836e478b4624789"): "LINK", // Chainlink LINK test (fréquent sur Uniswap Sepolia)
	common.HexToAddress("0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984"): "UNI",  // UNI (gouvernance Uniswap, déployé sur testnets)
	common.HexToAddress("0x5bb220afc6e2e008cb2302a83536a019ed245aa2"): "AAVE", // AAVE test token

	// --- AUTRES UTILES SUR SEPOLIA (souvent vus dans pools Uniswap V3) ---
	// Ajoute selon tes besoins, ex. :
	// common.HexToAddress("0x..."): "PYUSD", // Pas standard sur Sepolia
	// common.HexToAddress("0x..."): "LDO",  // Rare
}
