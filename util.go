package main

import (
	"context"
	"crypto/ecdsa"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jeffprestes/goethereumhelper"
)

func getAccountDetailsRinkeby(client *ethclient.Client) (privateKey *ecdsa.PrivateKey, publicKeyECDSA *ecdsa.PublicKey, fromAddress common.Address, nonce uint64, err error) {
	privateKey, err = crypto.HexToECDSA("e452c91efead165ec9ba005b5f437be6101eb10a7c650a3d0cc5da80d9ed5d80")
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
