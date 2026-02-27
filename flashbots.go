package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/metachris/flashbotsrpc"
)

func SendToFlashbots(client *ethclient.Client, bundleTxs []string) {
	// 1. Initialisation du client Flashbots
	// On utilise une clé privée pour s'identifier auprès du relai
	fb := flashbotsrpc.New("https://relay.flashbots.net")

	// 2. On récupère le bloc cible (souvent le bloc actuel + 1)
	blockNumber, _ := client.BlockNumber(context.Background())
	targetBlock := hexutil.EncodeUint64(blockNumber + 1)

	// 3. Préparation des arguments du Bundle
	args := flashbotsrpc.FlashbotsSendBundleArgs{
		Txs:         bundleTxs,   // Liste de tes TXs en RawHex
		BlockNumber: targetBlock, // Le bundle n'est valide que pour ce bloc
	}

	// 4. Envoi
	// relaySigningKey est une clé privée Go standard (*ecdsa.PrivateKey)
	res, err := fb.SendBundle(relaySigningKey, args)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Bundle Hash:", res.BundleHash)
}
