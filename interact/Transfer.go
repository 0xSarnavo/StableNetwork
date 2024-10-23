package main

import (
    "context"
    "crypto/ecdsa"
    "fmt"
    "log"
    "math/big"
	"os"

    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("/Users/darklord/Developer/Solidity-by-Projects/StableNetwork/.env")
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

    client, err := ethclient.Dial(os.Getenv("RPC"))
    if err != nil {
        log.Fatal("Failed to connect to RPC: ", err)
    }

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

    nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
    if err != nil {
        log.Fatal(err)
    }

    value := big.NewInt(1000000000000000000) 
    gasLimit := uint64(21000)                
    gasPrice, err := client.SuggestGasPrice(context.Background())
    if err != nil {
        log.Fatal(err)
    }

    toAddress := common.HexToAddress("0x4A1Db27cB7FAd2D842ca67A0b0C8BafBF3ca862a")
    var data []byte
    tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

    chainID, err := client.NetworkID(context.Background())
    if err != nil {
        log.Fatal(err)
    }

    signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
    if err != nil {
        log.Fatal(err)
    }

    err = client.SendTransaction(context.Background(), signedTx)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}