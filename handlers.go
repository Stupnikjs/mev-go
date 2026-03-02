package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

var SwapRouterV3Targets = map[string]string{
	"1f0464d1": "multicall",
	"04e45aaf": "exactInputSingle",
	"b858183f": "exactInput",
	"5023b4df": "exactOutputSingle",
	"09b81346": "exactOutput",
	"472b43f3": "swapExactTokensForTokens",
}

var SwapRouterV2Targets = map[string]string{
	// --- Swaps de Tokens (Montant d'entrée fixe) ---
	"38ed1739": "swapExactTokensForTokens",
	"7ff36ab5": "swapExactETHForTokens",
	"18cbafe5": "swapExactTokensForETH", // ⚠️ Parfois utilisé selon la version du compilateur

	// --- Swaps de Tokens (Montant de sortie fixe) ---
	"8803dbee": "swapTokensForExactTokens",
	"fb3bdb41": "swapETHForExactTokens",
	"4a25d94a": "swapTokensForExactETH",

	// --- Fee on Transfer (Tokens à taxe) ---
	"5c11d795": "swapExactTokensForTokensSupportingFeeOnTransferTokens",
	"b6f9de95": "swapExactETHForTokensSupportingFeeOnTransferTokens",
	"791ac947": "swapExactTokensForETHSupportingFeeOnTransferTokens",

	// --- Liquidité ---
	"e8e33700": "addLiquidity",
	"f305d719": "addLiquidityETH",
	"baa2abde": "removeLiquidity",
	"02751cec": "removeLiquidityETH",
	"af297903": "removeLiquidityETHSupportingFeeOnTransferTokens",
}

type AbiCache map[string]*abi.ABI

type Cache struct {
	m AbiCache
}

func (c *Cache) LoadCache() error {
	c.m = make(AbiCache)
	curr, err := os.Getwd()
	p := path.Join(curr, "abi", "uni", "v3", "swaprouter02.json")
	routerv3Abi, err := loadAbi(p)
	if err != nil {
		return err
	}
	for k := range SwapRouterV3Targets {
		c.m[k] = routerv3Abi
	}
	p = path.Join(curr, "abi", "uni", "v2", "swaprouter02.json")
	routerv2Abi, err := loadAbi(p)
	for k := range SwapRouterV2Targets {
		c.m[k] = routerv2Abi
	}
	return nil
}

func loadAbi(path string) (*abi.ABI, error) {
	data, err := os.ReadFile(path) // ← lit TOUT le fichier en mémoire
	if err != nil {
		return nil, fmt.Errorf("échec lecture %s: %w", path, err)
	}

	// Optionnel : vérification minimale que c'est du JSON
	var probe any
	if err := json.Unmarshal(data, &probe); err != nil {
		return nil, fmt.Errorf("fichier %s n'est pas du JSON valide: %w", path, err)
	}

	var abiItems any
	err = json.Unmarshal(data, &abiItems)
	if err == nil {
		// C'est un tableau → on le re-marshal en string et on parse
		abiJSON, err := json.Marshal(abiItems)
		if err != nil {
			return nil, fmt.Errorf("échec re-marshal tableau ABI: %w", err)
		}
		parsed, err := abi.JSON(strings.NewReader(string(abiJSON)))
		// 1. Chercher par ID (On s'assure que l'ID est bien formaté)

		return &parsed, err

	}
	return nil, err
}

func (e *Extractor) ProcessCallData(data []byte) {

	methodID := data[:4]
	methodIDHex := hex.EncodeToString(methodID)
	fmt.Println(methodIDHex)
	if !IsInUniMaps(methodIDHex) {
		return
	}
	fmt.Println(SwapRouterV2Targets[methodIDHex])
	fmt.Println(SwapRouterV3Targets[methodIDHex])

	abiParsed := e.c.m[methodIDHex]
	meth, err := abiParsed.MethodById(methodID)
	if err != nil {
		fmt.Println(err)
	}

	args, err := meth.Inputs.Unpack(data[:4])
	fmt.Println(args...)
	fmt.Println(len(args))
	for _, a := range args {
		fmt.Println(a.(string))
	}

}

func IsInUniMaps(methodid string) bool {
	if SwapRouterV2Targets[methodid] != "" {
		return true
	} else {
		if SwapRouterV3Targets[methodid] != "" {
			return true
		}
		return false
	}

}
