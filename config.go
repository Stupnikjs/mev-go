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
