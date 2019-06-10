package main

import (
	"context"
	"log"
	"math/big"

	"bitbucket.org/janusplatform/abigen-example/crud"
	"github.com/ethereum/go-ethereum/common"
	"github.com/jeffprestes/goethereumhelper"
)

func testLocal() {
	auth, blockchain, pvtKey := goethereumhelper.GetMockBlockchain()
	_, accountAddr, err := goethereumhelper.GetPubKey(pvtKey)
	contractAddr, trx, contract, err := crud.DeployUserCrud(auth, blockchain)
	if err != nil {
		log.Fatalln("Nao foi possivel fazer o deploy do contrato. Erro: ", err.Error())
	}
	blockchain.Commit()
	log.Println("Endereco do contrato: ", contractAddr.String())
	log.Printf("Transacao: %+v\n", trx.Hash().String())

	addrNovoUsuario := common.HexToAddress("0x71c9Ab003bdC883287F9238A9c2e5A535222beEC")

	result, err := contract.IsUser(nil, addrNovoUsuario)
	if err != nil {
		log.Fatalln("Nao conseguiu chamar o contrato. Erro: ", err.Error())
	}
	if result {
		log.Println("Usuario cadastrado")
	} else {
		log.Println("Usuario nao cadastrado")
	}

	log.Println("")
	log.Println("Incluindo usuario...")
	nonce, err := blockchain.PendingNonceAt(context.Background(), accountAddr)
	if err != nil {
		log.Fatalln("Nao foi possivel obter o nonce da conta. Erro: ", err.Error())
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei (1 eth)
	auth.GasLimit = uint64(310000) // in units
	gasPrice, err := blockchain.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal("Nao foi possivel obter o preco do gas")
	}
	auth.GasPrice = gasPrice.Mul(gasPrice, big.NewInt(2))
	trx, err = contract.InsertUser(auth, addrNovoUsuario, "jeffprestes@gmail.com", big.NewInt(40))
	if err != nil {
		log.Fatalln("Nao foi possivel incluir um novo usuario. Erro: ", err.Error())
	}
	blockchain.Commit()
	log.Printf("Transacao: %+v\n", trx.Hash().String())

	result, err = contract.IsUser(nil, addrNovoUsuario)
	if err != nil {
		log.Fatalln("Nao conseguiu chamar o contrato. Erro: ", err.Error())
	}
	if result {
		log.Println("Usuario cadastrado")
	} else {
		log.Println("Usuario nao cadastrado")
	}
}
