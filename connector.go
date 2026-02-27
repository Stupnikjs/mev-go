package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

// Provider represents a single RPC connection
type Provider struct {
	Name   string
	Client *ethclient.Client
	IsWS   bool
}

// Extractor holds our collection of connections
type Extractor struct {
	Providers []*Provider
}

// NewExtractor initializes connections to multiple endpoints
func NewExtractor(endpoints map[string]string) (*Extractor, error) {
	var providers []*Provider

	for name, url := range endpoints {

		client, err := ethclient.Dial(url)
		if err != nil {
			log.Printf("Failed to connect to %s: %v", name, err)
			continue
		}

		// Check if it's a websocket connection for streaming
		isWS := (url[:2] == "ws")

		providers = append(providers, &Provider{
			Name:   name,
			Client: client,
			IsWS:   isWS,
		})
	}

	return &Extractor{Providers: providers}, nil
}

func (e *Extractor) TransactionDetails(hash common.Hash) {
	// Context with timeout is vital in MEV so one slow RPC doesn't block you
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Use your first available provider (or implement a switcher)
	client := e.Providers[0].Client

	tx, isPending, err := client.TransactionByHash(ctx, hash)
	if err != nil || !isPending || tx.To() == nil {
		return
	}

	if isPending {
		data := tx.Data()
		if len(data) >= 4 {
			methodID := data[:4] // Les 4 premiers octets = la fonction

			if METHODID[fmt.Sprintf("%x", methodID)] != "" {
				fmt.Println(METHODID[fmt.Sprintf("%x", methodID)])

			}

		}
	}
}

func (e *Extractor) ListenToMempool() {
	// 1. Pick a WebSocket provider from your struct
	var wsClient *rpc.Client
	for _, p := range e.Providers {
		if p.IsWS {
			wsClient = p.Client.Client() // Get underlying RPC client
			break
		}
	}

	if wsClient == nil {
		log.Fatal("No WebSocket provider available")
	}

	// 2. Create a channel for hashes
	txHashes := make(chan common.Hash)

	// 3. Use gethclient for specialized subscriptions
	gClient := gethclient.New(wsClient)
	sub, err := gClient.SubscribePendingTransactions(context.Background(), txHashes)
	if err != nil {
		log.Fatal("Mempool subscription failed:", err)
	}
	defer sub.Unsubscribe()

	fmt.Println("🚀 Monitoring mempool for new hashes...")

	for {
		select {
		case err := <-sub.Err():
			log.Printf("Subscription error: %v", err)
			return
		case hash := <-txHashes:
			// You've caught a fish!
			// Now you need to fetch the full transaction details.
			go e.TransactionDetails(hash)
		}
	}
}

func (e *Extractor) ExtractDeepData(params []byte) {
	// On scanne par blocs de 32 octets (format standard EVM)
	for i := 0; i <= len(params)-32; i += 32 {
		block := params[i : i+32]

		// Les adresses ont 12 octets de zéros au début (padding)
		// [000000000000][Adresse sur 20 octets]
		potentialAddr := common.BytesToAddress(block[12:])

		if potentialAddr != (common.Address{}) {
			// Tu as trouvé un token ou un pool !
			fmt.Printf("     📍 Adresse trouvée : %s\n", potentialAddr.Hex())
		}
	}
}
