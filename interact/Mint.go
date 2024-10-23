package main

import (
    "context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"
	//"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
    NativeCoinMint "stable/contracts"
)

func main() {
    err := godotenv.Load("/Users/darklord/Developer/Solidity-by-Projects/StableNetwork/.env")
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	// Connect to the Ethereum client
	client, err := ethclient.Dial(os.Getenv("RPC"))
	if err != nil {
		log.Fatal("Failed to connect to RPC: ", err)
	}

	// Load the private key
	privateKey, err := crypto.HexToECDSA(os.Getenv("PrivateKey"))
	if err != nil {
		log.Fatal("Invalid private key: ", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Failed to assert public key type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// Retrieve the nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal("Failed to retrieve nonce: ", err)
	}

	// Suggest gas price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal("Failed to retrieve gas price: ", err)
	}

	// Parse the chain ID from the environment variable
	chainID := new(big.Int)
	if _, success := chainID.SetString(os.Getenv("ChainID"), 10); !success { // SetString returns *big.Int and a bool
		log.Fatal("Invalid Chain ID: ", os.Getenv("ChainID"))
	}

	// Create a transaction signer
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatalf("Failed to create transaction signer: %v", err)
	}

	// Set transaction parameters
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // Set to zero if not sending Ether
	auth.GasLimit = uint64(3000000)
	auth.GasPrice = gasPrice

    address := common.HexToAddress(os.Getenv("ContractAddress"))
	fmt.Println(address)
    instance, err := NativeCoinMint.NewNativeCoinMint(address, client)
    if err != nil {
        log.Fatal(err)
    }

	addr := common.HexToAddress("0x4A1Db27cB7FAd2D842ca67A0b0C8BafBF3ca862a")
	value := big.NewInt(10)
    tx, err := instance.Mint(auth, addr, value)
	if err != nil {
  		log.Fatal("Failed to mint: ", err)
	}
	fmt.Printf("Transaction sent: %s\n", tx.Hash().Hex())
}