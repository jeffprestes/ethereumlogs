package main

import (
	"context"
	"crypto/ecdsa"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jeffprestes/goethereumhelper"
)

func getAccountDetailsRinkeby(client *ethclient.Client) (privateKey *ecdsa.PrivateKey, publicKeyECDSA *ecdsa.PublicKey, fromAddress common.Address, nonce uint64, err error) {
	privateKey, err = crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		log.Fatal(err)
		return
	}

	_, fromAddress, err = goethereumhelper.GetPubKey(privateKey)
	if err != nil {
		log.Fatal(err)
		return
	}
	nonce, err = client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatalln(err)
		return
	}
	return
}
