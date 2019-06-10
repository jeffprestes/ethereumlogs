package main

import (
	"context"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/jeffprestes/goethereumhelper"
)

func sendEther() {
	log.Println("")
	log.Println("Conectando na Rinkeby...")
	client, err := goethereumhelper.GetRinkebyClient()
	if err != nil {
		log.Fatalln("Nao foi possivel conectar a rinkeby. Erro: ", err.Error())
	}
	defer client.Close()
	log.Println("Conectado na Rinkeby...")

	privateKey, pubKey, addr, nonce, err := getAccountDetailsRinkeby(client)

	log.Println("Detalhes da conta...", pubKey, addr.String())

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalln("Nao foi possivel obter o gas price. Erro: ", err.Error())
	}
	gasPrice = gasPrice.Mul(gasPrice, big.NewInt(2))
	gasLimit := uint64(910000)

	valueToBeSent := big.NewInt(30000000000000000)

	toAddress := common.HexToAddress("0x71c9Ab003bdC883287F9238A9c2e5A535222beEC")
	tx := types.NewTransaction(nonce, toAddress, valueToBeSent, gasLimit, gasPrice, nil)

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

	log.Printf("tx sent: [ %s ]\n\n", signedTx.Hash().Hex())

	receipt, err := client.TransactionReceipt(context.Background(), signedTx.Hash())
	for err != nil {
		<-time.After(2 * time.Second)
		log.Printf("%s %s\n", signedTx.Hash().Hex(), err.Error())
		receipt, err = client.TransactionReceipt(context.Background(), signedTx.Hash())
	}
	// if err != nil {
	// 	log.Fatalln("TransactionReceipt ", err)
	// 	return
	// }
	log.Println("tx receipt status: ", receipt.Status)
	log.Println(receipt.BlockHash.Hex(), " ", receipt.BlockNumber, " ", receipt.GasUsed, " ", receipt.Size(), " ", receipt.TransactionIndex, " ", string(receipt.PostState))
	log.Println("Logs length ", len(receipt.Logs))
	for _, loggy := range receipt.Logs {
		log.Println("Logs ", string(loggy.Data))
	}

	block, err := client.BlockByNumber(context.Background(), receipt.BlockNumber)
	if err != nil {
		log.Fatalln("Erro pegando block: ", err)
		return
	}
	log.Println("Tempo do bloco: ", block.Time())
	log.Println("Dificuldade ", block.Difficulty().Uint64())
	log.Println("Numero de trsnsacoes: ", len(block.Transactions()))
	for _, tx := range block.Transactions() {
		log.Println("Transaction ", tx.Hash().Hex())
		log.Println(tx.Value().String())         // 10000000000000000
		log.Println(tx.Gas())                    // 105000
		log.Println(tx.GasPrice().Uint64())      // 102000000000
		log.Println("Nonce ", tx.Nonce())        // 110644
		log.Println("Dados ", string(tx.Data())) // []
		if tx.To() != nil {
			log.Println("Destinatario: ", tx.To().Hex())
		}
		log.Printf("\n\n\n")
	}

	log.Println("Conectando na Rinkeby via Websocket...")
	clientWs, err := goethereumhelper.GetRinkebyClientWebsocket()
	if err != nil {
		log.Fatal("Erro ao conectar via websocket ao no ", err, "\n")
		return
	}
	log.Println("Conectado na Rinkeby via Websocket...")
	defer clientWs.Close()

	query := ethereum.FilterQuery{
		Addresses: []common.Address{toAddress},
	}

	logs := make(chan types.Log)
	sub, err := clientWs.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal("filter logs ", err, "\n")
		return
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatalln("Erro nx tx ", err)
		case vLog := <-logs:
			log.Printf("Tx Log: %+v\n", vLog)
		}
	}

}
