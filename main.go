package main

import "fmt"

func main() {
	// "other": "wss://sepolia.infura.io/ws/v3/e587127983764e6284261ebf6b4aaedf"
	endpoints := map[string]string{"main": "wss://mainnet.infura.io/ws/v3/e587127983764e6284261ebf6b4aaedf"}

	extractor, err := NewExtractor(endpoints)
	if err != nil {
		fmt.Println(err)
	}
	extractor.ListenToMempool()

}
