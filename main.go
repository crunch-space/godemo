package main

import (
	"github.com/crunch-space/contract/crunchProtocol"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {

}

var ProtocolFilterer, _ = crunchProtocol.NewCrunchProtocolFilterer(common.Address{}, &ethclient.Client{})

func ParseLog(log types.Log) (event *crunchProtocol.CrunchProtocolDeployCrunchVendor, err error) {
	//解析event
	return ProtocolFilterer.ParseDeployCrunchVendor(log)
}
