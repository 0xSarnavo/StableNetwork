package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// Prompt user for RPC URL
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the RPC URL: ")
	rpcURL, _ := reader.ReadString('\n')
	rpcURL = rpcURL[:len(rpcURL)-1] // Remove the newline character

	// Connect to the Ethereum client using the provided RPC URL
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatal("Failed to connect to RPC: ", err)
	}

	fmt.Println("Connected to:", rpcURL)

	// Start from block 0 and track the latest processed block number
	latestBlockNumber := big.NewInt(0)

	// Infinite loop to keep checking for new blocks
	for {
		// Get the latest block header
		latestHeader, err := client.HeaderByNumber(context.Background(), nil)
		if err != nil {
			log.Fatal(err)
		}

		currentBlockNumber := latestHeader.Number

		// Process all blocks from the last processed block to the latest
		for i := new(big.Int).Set(latestBlockNumber); i.Cmp(currentBlockNumber) <= 0; i.Add(i, big.NewInt(1)) {
			block, err := client.BlockByNumber(context.Background(), i)
			if err != nil {
				log.Printf("Failed to retrieve block number %d: %v", i.Int64(), err)
				continue
			}

			// Display block details
			fmt.Println("==========================================================")
			fmt.Printf("Block Number: %d\n", block.Number().Uint64())
			fmt.Printf("Block Hash: %s\n", block.Hash().Hex())
			fmt.Printf("Block Time: %d\n", block.Time())
			fmt.Printf("Block Difficulty: %d\n", block.Difficulty().Uint64())
			fmt.Printf("Number of Transactions: %d\n", len(block.Transactions()))
			fmt.Println("==========================================================")

			// Loop through all transactions in the block
			for j, tx := range block.Transactions() {
				fmt.Printf("\n  Transaction %d\n", j+1)
				fmt.Printf("  ------------------------------\n")
				fmt.Printf("  Transaction Hash: %s\n", tx.Hash().Hex())
				fmt.Printf("  Value: %s Ether\n", tx.Value().String())
				fmt.Printf("  Gas: %d\n", tx.Gas())
				fmt.Printf("  Gas Price: %s\n", tx.GasPrice().String())
				fmt.Printf("  Nonce: %d\n", tx.Nonce())

				// If it's a contract creation transaction
				if tx.To() == nil {
					fmt.Println("  To: Contract Creation")
				} else {
					fmt.Printf("  To: %s\n", tx.To().Hex())
				}

				// Get the receipt of the transaction to check the status
				receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
				if err != nil {
					log.Printf("  Failed to retrieve receipt for transaction %s: %v", tx.Hash().Hex(), err)
					continue
				}

				fmt.Printf("  Transaction Status: %d (1 = Success, 0 = Failure)\n", receipt.Status)

				// Print all the logs from the transaction
				if len(receipt.Logs) > 0 {
					fmt.Println("  Logs:")
					for _, logEntry := range receipt.Logs {
						fmt.Printf("    Log Address: %s\n", logEntry.Address.Hex())
						fmt.Printf("    Log Topics: %v\n", logEntry.Topics)
						fmt.Printf("    Log Data: %x\n", logEntry.Data)
					}
				} else {
					fmt.Println("  No Logs")
				}

				fmt.Println("  ------------------------------\n")
			}

			fmt.Println("\n==========================================================\n")
		}

		// Update the latest processed block number to the current latest block
		latestBlockNumber.Set(currentBlockNumber).Add(latestBlockNumber, big.NewInt(1))

		// Sleep for a while before checking for new blocks again
		time.Sleep(10 * time.Second)
	}
}
