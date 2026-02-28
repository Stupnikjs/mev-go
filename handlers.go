package main

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
)

var ABIRegistry = make(map[string]abi.ABI)

// --- UNISWAP UNIVERSAL ROUTER ---
func UniversalRouterExecute(data []byte) {}

// --- UNISWAP V2 / SUSHISWAP ---
func V2SwapExactTokensForTokens(data []byte) {
	Scanner(data)

}
func V2SwapExactEthForTokens(data []byte) {
	Scanner(data)
}
func V2SwapExactTokensForEth(data []byte) {
	Scanner(data)
}
func V2SwapSupportingFee(data []byte) {
	Scanner(data)
}

// --- UNISWAP V3 ---
func V3ExactInputSingle(data []byte)  {}
func V3ExactInput(data []byte)        {}
func V3ExactOutputSingle(data []byte) {}
func V3Multicall(data []byte)         {}

// --- LIQUIDATIONS ---
func AaveV3Liquidate(data []byte)   {}
func CompoundLiquidate(data []byte) {}

// --- AGGRÉGATEURS ---
func OneinchSwap(data []byte) {}
