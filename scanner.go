package main

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

func Scanner(data []byte) {
	if len(data) < 4+32*3 { // Min: Selector + 2 montants + 1 offset
		return
	}

	// 1. Identifier la fonction (Selector)
	methodID := fmt.Sprintf("%x", data[:4])

	// 2. Trouver l'Offset du tableau "path"
	// Dans la plupart des swaps V2 (ExactTokensForTokens, ExactETHForTokens),
	// l'offset du tableau 'path' est au 3ème slot (index 68) ou 2ème slot.
	// Mais pour faire simple et rapide, on cherche le pointeur de tableau :

	var pathOffset int
	switch methodID {
	case "38ed1739", "18cbafe5", "5c11d795": // swapExactTokensForTokens, ForETH, SupportingFee
		pathOffset = int(new(big.Int).SetBytes(data[68:100]).Int64()) + 4
	case "7ff36ab5": // swapExactETHForTokens
		pathOffset = int(new(big.Int).SetBytes(data[36:68]).Int64()) + 4
	default:
		return // Pas une méthode V2 connue
	}

	// 3. Lire la longueur du tableau (le slot à l'offset trouvé)
	if len(data) < pathOffset+32 {
		return
	}
	pathLen := int(new(big.Int).SetBytes(data[pathOffset : pathOffset+32]).Int64())

	// 4. Scanner uniquement les adresses du Path
	fmt.Printf("🔍 Analyse du chemin (%d tokens)...\n", pathLen)

	for j := range pathLen {
		start := pathOffset + 32 + (j * 32)
		if len(data) < start+32 {
			break
		}

		// On prend les 20 derniers octets du bloc de 32 (padding EVM)
		addr := common.BytesToAddress(data[start+12 : start+32])

		if symbol, ok := ADDRESS_TO_SYMBOL[addr]; ok {
			fmt.Printf("   📍 [%d] %s (%s)\n", j, symbol, addr.Hex())
		} else {
			fmt.Printf("   ❓ [%d] Inconnu (%s)\n", j, addr.Hex())
		}
	}
}
