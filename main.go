package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

func main() {
	// {"main": "wss://mainnet.infura.io/ws/v3/e587127983764e6284261ebf6b4aaedf"}

	endpoints := map[string]string{"sepolia": "wss://sepolia.infura.io/ws/v3/e587127983764e6284261ebf6b4aaedf", "other": "wss://ethereum-sepolia-rpc.publicnode.com", "another": "wss://sepolia.gateway.tenderly.co"}

	extractor, err := NewExtractor(endpoints)
	if err != nil {
		fmt.Println(err)
	}

	// extractor.ListenToMempool()
	// pool sepolia

	var SepoliaPools = []Pool{
		{
			Address: common.HexToAddress("0x9799b5edc1aa7d3fad350309b08df3f64914e244"),
			Tokens: [2]common.Address{
				common.HexToAddress("0x1c7D4B196Cb0C7B01d743Fbc6116a902379C7238"), // USDC (souvent utilisé)
				common.HexToAddress("0xfFf9976782d46CC05630D1f6eBAb18b2324d6B14"), // WETH officiel
			},
		},

		// Tu peux en ajouter d'autres si tu les crées ou en trouves via l'app Uniswap sur Sepolia
	}

	for {

		for _, p := range SepoliaPools {
			balances := p.GetPoolBalance(extractor)
			fmt.Println(balances)
		}

	}

}
