package main

import "fmt"

func main() {
	infura := "https://mainnet.infura.io/v3/e587127983764e6284261ebf6b4aaedf"
	// "infura": infura,
	endpoints := map[string]string{"infura": infura, "some": "wss://0xrpc.io/eth", "other": "https://ethereum-rpc.publicnode.com"}

	extractor, err := NewExtractor(endpoints)
	if err != nil {
		fmt.Println(err)
	}
	extractor.ListenToMempool()

}
