package main

import (
	"context"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (

	// ABI pour ERC20 balanceOf
	erc20ABI, _ = abi.JSON(strings.NewReader(`[{"inputs":[{"name":"account","type":"address"}],"name":"balanceOf","outputs":[{"name":"","type":"uint256"}],"stateMutability":"view","type":"function"}]`))
)

type Pool struct {
	Address common.Address
	Tokens  [2]common.Address
	Name    string
}

func (p *Pool) GetPoolBalance(extractor *Extractor) []*big.Int {

	ctx := context.Background()

	balanceToken0, err := getTokenBalance(ctx, extractor.Providers[0].Client, p.Tokens[0], p.Address)
	if err != nil {
		log.Printf("Erreur balance token0: %v", err)
	}

	if len(extractor.Providers) > 1 {
		balanceToken1, err := getTokenBalance(ctx, extractor.Providers[1].Client, p.Tokens[1], p.Address)
		if err != nil {
			log.Printf("Erreur balance token1: %v", err)
		}
		return []*big.Int{balanceToken0, balanceToken1}

	}

	balanceToken1, err := getTokenBalance(ctx, extractor.Providers[0].Client, p.Tokens[1], p.Address)
	if err != nil {
		log.Printf("Erreur balance token1: %v", err)
	}

	return []*big.Int{balanceToken0, balanceToken1}

}

// Appel balanceOf sur un token ERC20
func getTokenBalance(ctx context.Context, client *ethclient.Client, tokenAddr, account common.Address) (*big.Int, error) {
	data, err := erc20ABI.Pack("balanceOf", account)
	if err != nil {
		return nil, err
	}
	msg := ethereum.CallMsg{To: &tokenAddr, Data: data}
	result, err := client.CallContract(ctx, msg, nil)
	if err != nil {
		return nil, err
	}
	var balance *big.Int
	err = erc20ABI.UnpackIntoInterface(&balance, "balanceOf", result)
	return balance, err
}
