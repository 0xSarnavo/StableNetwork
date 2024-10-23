package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"math/big"
	"os"

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

	// Get the latest block header to find out the latest block number
	latestHeader, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	latestBlockNumber := latestHeader.Number
	fmt.Printf("The latest block number is: %s\n", latestBlockNumber.String())

	// Infinite loop to keep asking for block numbers and displaying their transactions
	for {
		// Prompt user for a block number to retrieve transactions from
		fmt.Print("Enter a block number to retrieve transactions from (or 'exit' to quit): ")
		inputBlockNumberStr, _ := reader.ReadString('\n')
		inputBlockNumberStr = inputBlockNumberStr[:len(inputBlockNumberStr)-1] // Remove the newline character

		// Exit if the user types "exit"
		if inputBlockNumberStr == "exit" {
			fmt.Println("Exiting the program.")
			break
		}

		// Convert the user's input to a big.Int for block number
		inputBlockNumber := new(big.Int)
		inputBlockNumber, success := inputBlockNumber.SetString(inputBlockNumberStr, 10)
		if !success {
			fmt.Println("Invalid block number entered. Please try again.")
			continue
		}

		// Retrieve and display the block
		block, err := client.BlockByNumber(context.Background(), inputBlockNumber)
		if err != nil {
			log.Printf("Failed to retrieve block number %d: %v", inputBlockNumber.Int64(), err)
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
}
