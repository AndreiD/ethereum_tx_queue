package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"time"
	"txqueue/database"
	"txqueue/eth"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
)

// Consume
func Consume() {

	maxRetries, _ := strconv.Atoi(os.Getenv("RETRY_COUNT"))
	retryTime, _ := strconv.Atoi(os.Getenv("RETRY_DELAY_MS"))
	successSleepTime, _ := strconv.Atoi(os.Getenv("SLEEP_AFTER_TX_MS"))

	for {
		size := database.Que.Size()
		if size == 0 {
			time.Sleep(1 * time.Second)
			continue
		}

		tx, err := database.Que.Pop()
		if err != nil {
			fmt.Println("error consuming queue:", err)
			continue
		}

		retries := 0
		for {
			fmt.Printf("processing item in queue: %d\n", size)
			hash, err := process(tx)
			if err == nil {
				fmt.Printf("tx processed successfully. tx id: %s", hash)
				appendToFile("success.log", string(tx)+","+hash)
				time.Sleep(time.Duration(successSleepTime) * time.Millisecond)
				break
			}

			fmt.Println("error processing tx:", err)

			retries++
			if retries > maxRetries {
				fmt.Println("max retry limit reached, abandoning tx")
				appendToFile("error.log", string(tx)+","+err.Error())
				break
			}

			fmt.Printf("retrying: %v\n", retries)
			time.Sleep(time.Duration(retryTime) * time.Millisecond)

		}
	}
}

// process sends the rawtx to the RPC
func process(rawTx []byte) (string, error) {
	rawTxBytes, err := hex.DecodeString(string(rawTx))
	if err != nil {
		return "", err
	}
	tx := new(types.Transaction)
	rlp.DecodeBytes(rawTxBytes, &tx)
	err = eth.Client.SendTransaction(context.Background(), tx)
	if err != nil {
		return "", err
	}
	return tx.Hash().Hex(), nil
}

// appendToErrorFile appends a string to a text file
func appendToFile(filename, data string) error {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(data + "\n")
	if err != nil {
		return err
	}
	return nil
}
