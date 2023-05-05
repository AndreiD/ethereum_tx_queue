package main

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

// https://testnet.snowtrace.io/address/0x721A8a1B6f313532c74e74C7e5Df3268f9B23917  - fell free to use it!
const CONTRACT_ADD = "0x721A8a1B6f313532c74e74C7e5Df3268f9B23917"

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}
	// Replace with your actual private key
	privateKeyHex := os.Getenv("PRIVATE_KEY")
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatal(err)
	}

	// Create an Ethereum client using the JSON-RPC URL
	rpcURL := os.Getenv("RPC_URL")
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}

	message := "hi #"
	for i := 0; i < 10; i++ {
		auth.Nonce = big.NewInt(int64(nonce))
		auth.Value = big.NewInt(0)
		auth.GasLimit = uint64(50000)
		auth.GasPrice = gasPrice
		auth.NoSend = true

		address := common.HexToAddress(CONTRACT_ADD)
		instance, err := NewHELLO(address, client)
		if err != nil {
			log.Fatalf("Failed to instantiate contract: %v", err)
		}

		tx, err := instance.SetMessage(auth, message+fmt.Sprintf("%d", i))
		if err != nil {
			log.Fatalf("Failed to send transaction: %v", err)
		}

		// Get the raw transaction data
		data, err := tx.MarshalBinary()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("sent one tx to the server")
		nonce += 1

		// Define the JSON payload
		payload := map[string]string{"rawTx": hex.EncodeToString(data)}
		jsonPayload, err := json.Marshal(payload)
		if err != nil {
			log.Fatal(err)
		}

		// Create the HTTP request
		req, err := http.NewRequest("POST", "http://localhost:5000/push", bytes.NewBuffer(jsonPayload))
		if err != nil {
			log.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")

		// Send the HTTP request
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		time.Sleep(2 * time.Second)

		// Print the response status code
		log.Printf("Response status code: %d\n", resp.StatusCode)
	}

}
