package main

import (
	"github.com/ethereum/go-ethereum/common"
)

var METHODID = map[string]func([]byte){
	// --- UNISWAP UNIVERSAL ROUTER ---
	"501d976c": UniversalRouterExecute,

	// --- UNISWAP V2 / SUSHISWAP ---
	"38ed1739": V2SwapExactTokensForTokens,
	"7ff36ab5": V2SwapExactEthForTokens,
	"18cbafe5": V2SwapExactTokensForEth,
	"5c11d795": V2SwapSupportingFee,

	// --- UNISWAP V3 ---
	"414bf389": V3ExactInputSingle,
	"c04b8d59": V3ExactInput,
	"db3e2135": V3ExactOutputSingle,
	"ac9650d8": V3Multicall,

	// --- LIQUIDATIONS ---
	"00a71129": AaveV3Liquidate,
	"fdb5a03e": CompoundLiquidate,

	// --- AGGRÉGATEURS ---
	"12a7b914": OneinchSwap, // "1inch" n'est pas un identifiant valide en Go, corrigé en "Oneinch"
}

// GetPools reserve

var ADDRESS_TO_SYMBOL = map[common.Address]string{
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
}
