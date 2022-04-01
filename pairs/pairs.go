package pairs

import (
	"fmt"
	"log"

	"github.com/Rask467/tj-arbitrage/joe"
	"github.com/Rask467/tj-arbitrage/pair"
	"github.com/ava-labs/coreth/ethclient"
	"github.com/ethereum/go-ethereum/common"
)

const (
	JOE_FACTORY      = "0x9Ad6C38BE94206cA50bb0d90783181662f0Cfa10"
	SUSHI_FACTORY    = "0xc35DADB65012eC5796536bD9864eD8773aBc74C4"
	PANGOLIN_FACTORY = "0xefa94DE7a4656D787667C749f7E1223D71E9FD88"

	AVAX      = "0xB31f66AA3C1e785363F0875A1B74E27b85FD66c7" // 18 decimals
	TEST_AVAX = "0xd00ae08403b9bbb9124bb305c09058e32c39a48c"
	USDT      = "0xc7198437980c041c805A1EDcbA50c1Ce5db95118" // 6 decimals
	TEST_USDT = "0xd8b917cf32022e35e09bac2c6f16a64fa7d1dec9"
)

func GetJoePair(rpcClient ethclient.Client) (*pair.Pair, common.Address) {
	joeFactoryAddr := common.HexToAddress(JOE_FACTORY)
	joeFactory, err := joe.NewJoe(joeFactoryAddr, rpcClient)
	if err != nil {
		log.Fatal("FAILED to create NewJoe: ", err)
	}

	token0Addr := common.HexToAddress(AVAX)
	token1Addr := common.HexToAddress(USDT)
	joePairAddr, err := joeFactory.GetPair(nil, token0Addr, token1Addr)
	if err != nil {
		log.Fatal("FAILED to GetPair joeFactory ", err)
	}
	fmt.Println("JOE PAIR ADDR: ", joePairAddr)

	joePair, err := pair.NewPair(joePairAddr, rpcClient)
	if err != nil {
		log.Fatal("JOE PAIR FAILED", err)
	}

	return joePair, joePairAddr
}

func GetSushiPair(rpcClient ethclient.Client) (*pair.Pair, common.Address) {
	sushiAddr := common.HexToAddress(SUSHI_FACTORY)
	sushiFactory, err := joe.NewJoe(sushiAddr, rpcClient)
	if err != nil {
		log.Fatal("FAILED NEW SUSHI", err)
	}

	token0Addr := common.HexToAddress(AVAX)
	token1Addr := common.HexToAddress(USDT)
	sushiPairAddr, err := sushiFactory.GetPair(nil, token0Addr, token1Addr)
	if err != nil {
		log.Fatal("FAILED GET PAIR SUSHI", err)
	}
	fmt.Println("SUSHI PAIR ADDR: ", sushiPairAddr)

	sushiPair, err := pair.NewPair(sushiPairAddr, rpcClient)
	if err != nil {
		log.Fatal("SUSHI NEW PAIR FAILED", err)
	}

	return sushiPair, sushiPairAddr
}

func GetPangolinPair(rpcClient ethclient.Client) (*pair.Pair, common.Address) {
	pangolinAddr := common.HexToAddress(PANGOLIN_FACTORY)
	pangolinFactory, err := joe.NewJoe(pangolinAddr, rpcClient)
	if err != nil {
		log.Fatal(err)
	}

	token0Addr := common.HexToAddress(AVAX)
	token1Addr := common.HexToAddress(USDT)
	pangolinPairAddr, err := pangolinFactory.GetPair(nil, token0Addr, token1Addr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("PANGOLIN PAIR ADDR: ", pangolinPairAddr)

	pangolinPair, err := pair.NewPair(pangolinPairAddr, rpcClient)
	if err != nil {
		log.Fatal(err)
	}

	return pangolinPair, pangolinPairAddr
}
