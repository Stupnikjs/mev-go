package main

var METHODID = map[string]string{
	// --- UNISWAP UNIVERSAL ROUTER ---
	"501d976c": "universal_router_execute",

	// --- UNISWAP V2 / SUSHISWAP ---
	"38ed1739": "v2_swap_exact_tokens_for_tokens",
	"7ff36ab5": "v2_swap_exact_eth_for_tokens",
	"18cbafe5": "v2_swap_exact_tokens_for_eth",
	"5c11d795": "v2_swap_supporting_fee",

	// --- UNISWAP V3 ---
	"414bf389": "v3_exact_input_single",
	"c04b8d59": "v3_exact_input",
	"db3e2135": "v3_exact_output_single",
	"ac9650d8": "v3_multicall",

	// --- LIQUIDATIONS ---
	"00a71129": "aave_v3_liquidate",
	"fdb5a03e": "compound_liquidate",

	// --- AGGRÉGATEURS ---
	"12a7b914": "1inch_swap",
}

// GetPools reserve

var TOKENS = map[string]string{
	// --- NATIVE / WRAPPED ---
	"WETH": "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2", // Wrapped Ether
	"WBTC": "0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C599", // Wrapped Bitcoin

	// --- STABLECOINS ---
	"USDC":  "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eb48", // USD Coin
	"USDT":  "0xdAC17F958D2ee523a2206206994597C13D831ec7", // Tether USD
	"DAI":   "0x6B175474E89094C44Da98b954EedeAC495271d0F", // Dai Stablecoin
	"EURA":  "0x1aBaEA1f7230f38fFB5A03b8395B4d3f1fd3872E", // Euro (Angle)
	"PYUSD": "0x6c3ea9036406852006290770bedfc107f186715f", // PayPal USD

	// --- BLUE CHIPS / DEFI ---
	"LINK": "0x514910771AF9Ca656af840dff83E8264EcF986CA", // Chainlink
	"UNI":  "0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984", // Uniswap
	"AAVE": "0x7Fc66500c84A76Ad7e9c93437bFc5Ac33E2DDaE9", // Aave
	"LDO":  "0x5A987829255094551674823AeeD7103027f156f8", // Lido DAO
	"MKR":  "0x9f8F72aA9304c8B593d555F12eF6589cC3A579A2", // Maker
	"COMP": "0xc00e94cb662c3520282e6f5717214004a7f26888", // Compound

	// --- LAYER 2 / BRIDGE ---
	"ARB": "0xB50721BCf8d664c30412Cfbc6cf7a15145234ad1", // Arbitrum (Mainnet bridge)
	"OP":  "0x4200000000000000000000000000000000000042", // Optimism
}
