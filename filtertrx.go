package main

import (
	"context"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/jeffprestes/goethereumhelper"
)

func FilterTrx() {
	log.Println("")
	log.Println("Conectando na Rinkeby...")
	client, err := goethereumhelper.GetRinkebyClient()
	if err != nil {
		log.Fatalln("Nao foi possivel conectar a rinkeby. Erro: ", err.Error())
	}
	defer client.Close()
	log.Println("Conectado na Rinkeby...")

	toAddress := common.HexToAddress("0x6fe0D8E112B8C4893e344f42e05d170503b559d3")

	// hashTest := common.HexToHash("0xee1504a83b6d4a361f4c1dc78ab59bfa30d6a3b6612c403e86bb01ef2984295f")
	// query := ethereum.FilterQuery{
	// 	Addresses: []common.Address{toAddress},
	// 	FromBlock: big.NewInt(4500403),
	// 	Topics:    [][]common.Hash{{hashTest}},
	// }

	query := ethereum.FilterQuery{
		Addresses: []common.Address{toAddress},
		FromBlock: big.NewInt(4500403),
	}

	log.Println("Buscando os logs...")
	logsTx, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatalln("Nao foi possivel conectar a rinkeby. Erro: ", err.Error())
	}

	fileReader, err := os.Open("atlas-token-abi.json")
	if err != nil {
		log.Fatalln("Nao foi possivel abrir o arquivo com ABI. Erro: ", err.Error())
	}

	log.Println("Carregando ABI para ler os logs...")
	//tokenAbi, err := abi.JSON(strings.NewReader(crud.UserCrudABI))
	tokenAbi, err := abi.JSON(fileReader)
	if err != nil {
		log.Fatalln("Nao foi possivel ler a ABI do contrato. Erro: ", err.Error())
	}

	log.Println("Eventos disponíveis...")
	for _, evento := range tokenAbi.Events {
		log.Println("   ", evento.String(), evento.Id().String())
	}

	log.Println("")
	log.Println("====================================")
	log.Println("Logs abaixo....")
	log.Println("")
	for _, logg := range logsTx {
		printLog(logg, tokenAbi)
	}

	log.Println("FIM")
}

func printLog(logg types.Log, abi abi.ABI) {
	if !logg.Removed && len(logg.Topics) > 0 {
		if len(logg.Topics) == 3 {
			logTransfers(logg, abi)
		} else if len(logg.Topics) == 2 {
			//event WhitelistedAdded
			//her hash is 0xee1504a83b6d4a361f4c1dc78ab59bfa30d6a3b6612c403e86bb01ef2984295f
			if logg.Topics[0].String() == "0xee1504a83b6d4a361f4c1dc78ab59bfa30d6a3b6612c403e86bb01ef2984295f" {
				logWhitelistAdd(logg)
			} else {
				// logNonIdentifiedTopic("two topics", logg, abi)
			}
		} else {
			//logNonIdentifiedTopic("different topic numbers", logg, abi)
		}
	}
}

func logNonIdentifiedTopic(prefix string, logg types.Log, abi abi.ABI) {
	logHeaders(logg)
	if len(logg.Topics) > 0 {
		topicTitle := logg.Topics[0]
		//log.Println("   ", topicTitle.String())
		for _, evento := range abi.Events {
			//log.Println("   ", topicTitle.String(), " - ", evento.Id().String())
			if topicTitle.String() == evento.Id().String() {
				log.Println("   ", prefix, " - ", evento.String(), evento.Id().String())
			}
		}
	}
	log.Println("====================================")
	log.Println("")
}

func logTransfers(logg types.Log, abi abi.ABI) {
	var transferEvent struct {
		From  common.Address
		To    common.Address
		Value *big.Int
	}

	err := abi.Unpack(&transferEvent, "Transfer", logg.Data)
	if err != nil {
		log.Println("Failed to unpack Transfer log")
		log.Println("====================================")
		log.Println("")
		return
	}

	logHeaders(logg)
	transferEvent.From = common.BytesToAddress(logg.Topics[1].Bytes())
	transferEvent.To = common.BytesToAddress(logg.Topics[2].Bytes())

	log.Println("Transferlog")
	log.Println("    From", transferEvent.From.Hex())
	log.Println("    To", transferEvent.To.Hex())
	log.Println("    Value", transferEvent.Value)
	log.Println("====================================")
	log.Println("")
}

func logWhitelistAdd(logg types.Log) {
	logHeaders(logg)

	var whitelistedAddedEvent struct {
		Account common.Address
	}
	whitelistedAddedEvent.Account = common.HexToAddress(logg.Topics[1].String())
	log.Println("     ", "A conta From", whitelistedAddedEvent.Account.Hex(), " foi adicionada ao Whitelist")
	log.Println("====================================")
	log.Println("")
}

func logHeaders(logg types.Log) {
	log.Println("Address: ", logg.Address.Hex())
	log.Println("BlockNumber: ", logg.BlockNumber)
	log.Println("TxHash: ", logg.TxHash.Hex())
	log.Println("TxIndex: ", logg.TxIndex)
	log.Println("Index: ", logg.Index)
	log.Println("Numero de topicos: ", len(logg.Topics))
}
