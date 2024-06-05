package main

import (
	"context"
	"log"

	"github.com/crunch-space/contract/crunchProtocol"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var wsUri = "wss://mainnet.infura.io/ws/v3/your_project_id"
var client, _ = ethclient.Dial(wsUri)

var ProtocolFilterer, _ = crunchProtocol.NewCrunchProtocolFilterer(common.Address{}, client)

// parse watch filter event
func ParseLog(log types.Log) (event *crunchProtocol.CrunchProtocolDeployCrunchVendor, err error) {
	return ProtocolFilterer.ParseDeployCrunchVendor(log)
}
func FilterLogs(blockNumber uint64) (err error) {
	var iterator *crunchProtocol.CrunchProtocolDeployCrunchVendorIterator
	iterator, err = ProtocolFilterer.FilterDeployCrunchVendor(&bind.FilterOpts{
		Start:   blockNumber,
		End:     &blockNumber,
		Context: context.Background(),
	}, nil, nil)
	if err != nil {
		return
	}
	for iterator.Next() {
		log.Println("FilterLogs", iterator.Event)
	}
	return iterator.Error()
}

func main() {
	log.Println("BlockSubscribe Start")
	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal("BlockSubscribe.Start.err", err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal("BlockSubscribe.err", err)
		case header := <-headers:
			//fmt.Println(header.Hash().Hex()) // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f
			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal("BlockSubscribe.headers.err", err)
			}
			var blockNumber = block.NumberU64()
			log.Println("BlockNumber", blockNumber)
		}
	}
}
