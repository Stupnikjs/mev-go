package main

var METHODID = map[string]func([]byte){
	// --- UNISWAP UNIVERSAL ROUTER ---
	"501d976c": UniversalRouterExecute,

	// --- UNISWAP V3 ---
	"414bf389": V3ExactInputSingle,
	"c04b8d59": V3ExactInput,
	"db3e2135": V3ExactOutputSingle,
	"ac9650d8": V3Multicall,
}

// --- UNISWAP UNIVERSAL ROUTER ---
func UniversalRouterExecute(data []byte) {}

// --- UNISWAP V3 ---
func V3ExactInputSingle(data []byte)  { ParseSwapCalldata(data) }
func V3ExactInput(data []byte)        { ParseSwapCalldata(data) }
func V3ExactOutputSingle(data []byte) { ParseSwapCalldata(data) }
func V3Multicall(data []byte)         { ParseSwapCalldata(data) }

// --- LIQUIDATIONS ---
func AaveV3Liquidate(data []byte)   {}
func CompoundLiquidate(data []byte) {}

// --- AGGRÉGATEURS ---
func OneinchSwap(data []byte) {}
