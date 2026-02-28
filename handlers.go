package main

import (
	"bytes"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

// --- UNISWAP UNIVERSAL ROUTER ---
func UniversalRouterExecute(data []byte) {}

// --- UNISWAP V2 / SUSHISWAP ---
func V2SwapExactTokensForTokens(data []byte) {
	// 1. Parser l'ABI
	r := bytes.NewReader(data)
	parsedABI, _ := abi.JSON(r)

	// 2. Créer une map pour stocker les arguments
	params := make(map[string]interface{})

	// 3. Unpack (on ignore les 4 premiers octets du MethodID)
	err := parsedABI.Methods["swapExactTokensForTokens"].Inputs.UnpackIntoMap(params, data)
	if err != nil {
		return
	}

}
func V2SwapExactEthForTokens(data []byte) {
	// 1. Parser l'ABI
	r := bytes.NewReader(data)
	parsedABI, _ := abi.JSON(r)

	fmt.Println(parsedABI.Methods)
	fmt.Println(parsedABI)
}
func V2SwapExactTokensForEth(data []byte) {}
func V2SwapSupportingFee(data []byte)     {}

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
