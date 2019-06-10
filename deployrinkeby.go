package main

import (
	"context"
	"log"
	"math/big"

	"bitbucket.org/janusplatform/abigen-example/crud"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/jeffprestes/goethereumhelper"
)

func testRinkeby() {
	log.Println("")
	log.Println("Conectando na Rinkeby...")
	client, err := goethereumhelper.GetRinkebyClient()
	if err != nil {
		log.Fatalln("Nao foi possivel conectar a rinkeby. Erro: ", err.Error())
	}
	privateKey, pubKey, addr, nonce, err := getAccountDetailsRinkeby(client)

	log.Println("Detalhes da conta...", pubKey, addr.String())

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	//auth.GasPrice
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalln("Nao foi possivel obter o gas price. Erro: ", err.Error())
	}
	auth.GasPrice = gasPrice.Mul(gasPrice, big.NewInt(2))
	auth.GasLimit = 910000

	log.Println("Fazendo deploy na Rinkeby...")
	contractAddr, trx, contractInstance, err := crud.DeployUserCrud(auth, client)
	if err != nil {
		log.Fatalln("Nao foi possivel fazer o deploy do contrato. Erro: ", err.Error())
	}
	log.Println("Endereco do contrato: ", contractAddr.String())
	log.Printf("Transacao: %+v\n", trx.Hash().String())
	log.Printf("Instancia contrato: %+v\n", contractInstance)

	clientWebSocket, err := goethereumhelper.GetRinkebyClientWebsocket()
	if err != nil {
		log.Fatalln("Nao foi possivel conectar a rinkeby via websocket. Erro: ", err.Error())
	}
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddr},
	}

	logs := make(chan types.Log)
	sub, err := clientWebSocket.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatalln("Nao foi possivel subscrever ao evento de deploy do contrato. Erro: ", err.Error())
	}

	log.Println("Aguardando a transacao ser confirmada na rede...")

	crud.NewUserCrudFilterer(contractAddr, nil)

	for {
		select {
		case err := <-sub.Err():
			log.Fatalln("Erro ao escutar os logs. Erro: ", err.Error())
		case vLog := <-logs:
			log.Println(vLog)
		}
	}
}
