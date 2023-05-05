package eth

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Client the eth client
var Client *ethclient.Client

// InitEthClient initialises the the client
// noinspection GoNilness
func InitEthClient(url string) error {
	ethClient, err := ethclient.Dial(url)
	if err != nil {
		return err
	}

	fmt.Printf("\nconnected to the ETH provider %s\n", url)

	// is Syncing ?
	isSyncing, err := IsSyncying(ethClient)
	if err != nil {
		fmt.Errorf("can't get eth client status %s", err.Error())
	}
	if isSyncing {
		fmt.Errorf("!!! The ETH client is Syncing !!!")
	}

	// block number
	blockNumber, err := GetBlockNumber(ethClient)
	if err != nil {
		fmt.Errorf("can't get current block number %s", err.Error())
	}
	fmt.Printf("\ncurrent block number %s\n", blockNumber)

	Client = ethClient

	return nil
}

// GetBlockNumber gets the block number
func GetBlockNumber(client *ethclient.Client) (string, error) {
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return "", err
	}
	return header.Number.String(), nil
}

// GetTX gets the tx by hash
func GetTX(client *ethclient.Client, txID string) (*types.Transaction, bool, error) {
	tx, isPending, err := client.TransactionByHash(context.Background(), common.HexToHash(txID))
	if err != nil {
		return nil, false, err
	}
	return tx, isPending, nil
}

// GetTXReceipt gets the tx by hash
func GetTXReceipt(client *ethclient.Client, txID string) (*types.Receipt, error) {
	receipt, err := client.TransactionReceipt(context.Background(), common.HexToHash(txID))
	if err != nil {
		return nil, err
	}
	return receipt, nil
}

// IsSyncying returns false if it's not syncing
func IsSyncying(client *ethclient.Client) (bool, error) {
	sync, err := client.SyncProgress(context.Background())
	if err != nil {
		return false, err
	}
	if sync == nil {
		return false, nil
	}
	return true, nil
}
