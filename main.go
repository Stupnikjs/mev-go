package main

import "fmt"

func main() {
	// infura := "https://mainnet.infura.io/v3/e587127983764e6284261ebf6b4aaedf"
	// "infura": infura,
	endpoints := map[string]string{"other": "wss://lb.drpc.live/ethereum/AhuxMhCqfkI8pF_0y4Fpi89GWcIMFIwR8ZsatuZZzRRv"}

	extractor, err := NewExtractor(endpoints)
	if err != nil {
		fmt.Println(err)
	}
	extractor.ListenToMempool()

}
