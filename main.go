package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/Rask467/tj-arbitrage/pair"
	"github.com/Rask467/tj-arbitrage/pairs"
	"github.com/ava-labs/coreth/accounts/abi"
	"github.com/ava-labs/coreth/accounts/abi/bind"
	"github.com/ava-labs/coreth/core/types"
	"github.com/ava-labs/coreth/ethclient"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/joho/godotenv"
)

// DEPLOYED ADDRESS = 0x9208c8cBF9841Cb8de3fFb06082bdD91c8A92311
// DEPLOYED ADDRESS WITHOUT PROFIT CONSTRAINT = 0xf4E17b8F3951F8bda7703FA3880a5130507522BC

const (
	E_SUSHI_FACTORY    = "0xC0AEe478e3658e2610c5F7A4A2E1777cE9e4f2Ac"
	UNI_FACTORY        = "0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f"
	WETH               = "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"
	E_USDT             = "0xdAC17F958D2ee523a2206206994597C13D831ec7"
	SC_ADDR            = "0xf4E17b8F3951F8bda7703FA3880a5130507522BC"
	JOE_AVAX_USDT      = "0xeD8CBD9F0cE3C6986b22002F03c6475CEb7a6256"
	SUSHI_AVAX_USDT    = "0x09657b445dF5BF0141e3EF0f5276a329fc01DE01"
	PANGOLIN_AVAX_USDT = "0xe28984e1EE8D431346D32BeC9Ec800Efb643eef4"
)

type Reserve struct {
	Address common.Address
	R0      *big.Int
	R1      *big.Int
}

type Market struct {
	Address  common.Address
	Pair     *pair.Pair
	PriceIn  *big.Int
	PriceOut *big.Int
	Reserve  Reserve
}

type CrossedMarket struct {
	BuyMarket  Market
	SellMarket Market
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// https://api.avax-test.network/ext/bc/C/rpc
	// https://api.avax.network/ext/bc/C/rpc
	rpcClient, err := ethclient.Dial("https://api.avax.network/ext/bc/C/rpc")
	if err != nil {
		log.Fatal(err)
	}

	//wss://eth-mainnet.alchemyapi.io/v2/zHWlgQu0BiHtPXsquJuOzgJ9eU7iWI6M
	//wss://api.avax-test.network/ext/bc/C/ws
	wsClient, err := ethclient.Dial("wss://api.avax.network/ext/bc/C/ws")
	if err != nil {
		log.Fatal(err)
	}
	headers := make(chan *types.Header)
	sub, err := wsClient.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(os.Getenv("PK"))
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Println("FROM ADDRESS: ", fromAddress)

	// Setup up pairs.
	joePair, joePairAddr := pairs.GetJoePair(rpcClient)
	sushiPair, sushiPairAddr := pairs.GetSushiPair(rpcClient)
	pangolinPair, pangolinPairAddr := pairs.GetPangolinPair(rpcClient)

	joeToken0, err := joePair.Token0(nil)
	if err != nil {
		log.Fatal("JOE TOKEN 0 FAIL", err)
	}
	fmt.Println("joePair Token0: ", joeToken0)

	joeToken1, err := joePair.Token1(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("joePair Token1: ", joeToken1)

	sushiToken0, err := sushiPair.Token0(nil)
	if err != nil {
		log.Fatal("SUSHI TOKEN 0 FAIL", err)
	}
	fmt.Println("sushiPair Token0: ", sushiToken0)

	sushiToken1, err := sushiPair.Token1(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("sushiPair Token1: ", sushiToken1)

	pangolinToken0, err := pangolinPair.Token0(nil)
	if err != nil {
		log.Fatal("SUSHI TOKEN 0 FAIL", err)
	}
	fmt.Println("pangolinPair Token0: ", pangolinToken0)

	pangolinToken1, err := pangolinPair.Token1(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("pangolinPair Token1: ", pangolinToken1)

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			fmt.Println("BLOCK NUMBER: ", header.Number)
			fmt.Println("HEADER PARENT HASH: ", header.ParentHash)
			joeReserves, err := joePair.GetReserves(nil)
			if err != nil {
				log.Fatal("JOE RESERVES FAIL", err)
			}
			fmt.Println("JOE RESERVES: ", joeReserves)

			sushiReserves, err := sushiPair.GetReserves(nil)
			if err != nil {
				log.Fatal("SUSHI RESERVES FAIL", err)
			}
			fmt.Println("SUSHI RESERVES: ", sushiReserves)

			pangolinReserves, err := pangolinPair.GetReserves(nil)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("PANGOLIN RESERVES: ", pangolinReserves)

			reserves := map[*pair.Pair]Reserve{
				joePair: {
					Address: joePairAddr,
					R0:      joeReserves.Reserve0,
					R1:      joeReserves.Reserve1,
				},
				sushiPair: {
					Address: sushiPairAddr,
					R0:      sushiReserves.Reserve0,
					R1:      sushiReserves.Reserve1,
				},
				pangolinPair: {
					Address: pangolinPairAddr,
					R0:      pangolinReserves.Reserve0,
					R1:      pangolinReserves.Reserve1,
				},
			}

			var pricedMarkets []Market
			oneAvax := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)
			oneUSDT := new(big.Int).Exp(big.NewInt(10), big.NewInt(6), nil)

			for pair, reserve := range reserves {
				// Amount of AVAX in required to get one USDT
				AVAXIn := getAmountIn(oneUSDT, reserve.R0, reserve.R1)
				fmt.Println("AVAX IN: ", AVAXIn)

				// Amount of AVAX out for one USDT
				AVAXOut := getAmountOut(oneUSDT, reserve.R1, reserve.R0)
				fmt.Println("AVAX OUT: ", AVAXOut)

				market := Market{
					Address:  reserve.Address,
					Pair:     pair,
					PriceIn:  AVAXIn,
					PriceOut: AVAXOut,
					Reserve:  reserve,
				}
				pricedMarkets = append(pricedMarkets, market)
			}

			var crossedMarkets []CrossedMarket
			for _, m1 := range pricedMarkets {
				for _, m2 := range pricedMarkets {
					if m1.PriceOut.Cmp(m2.PriceIn) > 0 {
						crossedMarkets = append(crossedMarkets, CrossedMarket{BuyMarket: m2, SellMarket: m1})
					}
				}
			}

			if len(crossedMarkets) > 0 {
				// For now just use the first crossed market.
				buyMarket := crossedMarkets[0].BuyMarket
				sellMarket := crossedMarkets[0].SellMarket
				fmt.Println("FOUND ONE!!!!", crossedMarkets)
				fmt.Println("BUY MARKET: ", buyMarket)
				fmt.Println("SELL MARKET: ", sellMarket)

				nonce, err := rpcClient.AcceptedNonceAt(context.Background(), fromAddress)
				if err != nil {
					log.Fatal(err)
				}

				gasPrice, err := rpcClient.SuggestGasPrice(context.Background())
				if err != nil {
					log.Fatal(err)
				}

				// MAIN: 43114
				// TEST: 43113
				auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(43114))
				if err != nil {
					log.Fatal(err)
				}
				auth.Nonce = big.NewInt(int64(nonce))
				// auth.Value = big.NewInt(0)     // in wei
				auth.GasLimit = uint64(600000)                            // in units
				auth.GasPrice = new(big.Int).Mul(gasPrice, big.NewInt(2)) // TODO: Mess with the gas price, right now just doubling it to ensure quick inclusion in a block.

				content, err := os.ReadFile("pair.abi")
				if err != nil {
					log.Fatal("failed to read file: ", err)
				}

				pairABI, err := abi.JSON(strings.NewReader(string(content)))
				if err != nil {
					log.Fatal("Failed to create abi from JSON string: ", err)
				}
				method := pairABI.Methods["swap"]

				fmt.Println("METHOD", method)

				addrT, err := abi.NewType("address", "address", nil)
				if err != nil {
					log.Fatal("Failed to create new type address: ", err)
				}
				uintT, err := abi.NewType("uint256", "uint256", nil)
				if err != nil {
					log.Fatal("Failed to create new type uint: ", err)
				}
				args := abi.Arguments{
					{
						Type: addrT,
					},
					{
						Type: addrT,
					},
					{
						Type: uintT,
					},
					{
						Type: uintT,
					},
				}

				// Find the optimal amount of avax to buy/sell.
				lowerBound := new(big.Int).Div(oneAvax, big.NewInt(100))
				upperBound := new(big.Int).Mul(oneAvax, big.NewInt(10000))

				maxVolume := new(big.Int).Mul(oneAvax, big.NewInt(10000))
				minVolume := new(big.Int).Div(oneAvax, big.NewInt(100))

				maxAttempts := 20
				attempts := 0
				bestSize := big.NewInt(0)
				previousProfit := big.NewInt(0)
				usdtOut := big.NewInt(0)
				avaxOut := big.NewInt(0)
				profit := big.NewInt(0)
				currentTestVolume := new(big.Int).Div(maxVolume, big.NewInt(2))
				for {
					if attempts > maxAttempts || currentTestVolume.Cmp(maxVolume) >= 0 || currentTestVolume.Cmp(minVolume) <= 0 {
						break
					}
					fmt.Println("TEST VOLUME: ", currentTestVolume)
					fmt.Println("upperBound: ", upperBound)
					fmt.Println("lowerBound: ", lowerBound)
					tUSDTOut := getAmountOut(currentTestVolume, buyMarket.Reserve.R0, buyMarket.Reserve.R1)
					fmt.Println("USDTOut: ", tUSDTOut)
					tAvaxOut := getAmountOut(tUSDTOut, sellMarket.Reserve.R1, sellMarket.Reserve.R0)
					fmt.Println("AvaxOut: ", tAvaxOut)
					profit = new(big.Int).Sub(tAvaxOut, currentTestVolume)
					fmt.Println("Previous Profit: ", previousProfit)
					fmt.Println("PROFIT: ", profit)
					if profit.Cmp(previousProfit) > 0 {
						fmt.Println("Trying next size up.")
						// Then try next size up.
						bestSize = currentTestVolume
						previousProfit = profit
						usdtOut = tUSDTOut
						avaxOut = tAvaxOut

						// Update test volume.
						lowerBound = currentTestVolume
						currentTestVolume = new(big.Int).Add(new(big.Int).Div(new(big.Int).Sub(upperBound, currentTestVolume), big.NewInt(2)), currentTestVolume)
						attempts = attempts + 1
					} else {
						fmt.Println("Trying next size down.")
						// Update test volume.
						upperBound = currentTestVolume
						currentTestVolume = new(big.Int).Sub(currentTestVolume, new(big.Int).Div(new(big.Int).Sub(currentTestVolume, lowerBound), big.NewInt(2)))
						attempts = attempts + 1
					}
				}

				minimumProfit := new(big.Int).Div(oneAvax, big.NewInt(80))
				fmt.Println("MINIMUM PROFIT", minimumProfit)

				if profit.Cmp(big.NewInt(0)) <= 0 || profit.Cmp(minimumProfit) < 0 {
					fmt.Println("Final calculated profit is: ", profit)
					fmt.Println("CANCELLING TRADE: Not Profittable.")
					continue
				}

				fmt.Println("BEST SIZE: ", bestSize)
				fmt.Println("PREVIOUS PROFIT: ", previousProfit)
				fmt.Println("USDT OUT: ", usdtOut)
				fmt.Println("AVAX OUT: ", avaxOut)

				calldata, err := args.Pack(buyMarket.Address, sellMarket.Address, avaxOut, bestSize)
				if err != nil {
					log.Fatal("Failed to pack arguments: ", err)
				}

				scAddr := common.HexToAddress(SC_ADDR)
				txn, err := buyMarket.Pair.Swap(auth, big.NewInt(0), usdtOut, scAddr, calldata)
				if err != nil {
					log.Fatal("Failed to call Swap method on buyMarket", err)
				}

				fmt.Println("TRANSACTION HASH: ", txn.Hash())
				os.Exit(1)
			}
		}
	}
}

func getAmountIn(amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) *big.Int {
	numerator := new(big.Int).Mul(reserveIn, amountOut)
	numerator = new(big.Int).Mul(numerator, big.NewInt(1000))
	denominator := new(big.Int).Sub(reserveOut, amountOut)
	denominator = new(big.Int).Mul(denominator, big.NewInt(997))
	amountIn := new(big.Int).Div(numerator, denominator)
	amountIn = new(big.Int).Add(amountIn, big.NewInt(1))
	return amountIn
}

func getAmountOut(amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) *big.Int {
	amountInWithFee := new(big.Int).Mul(amountIn, big.NewInt(997))
	numerator := new(big.Int).Mul(amountInWithFee, reserveOut)
	denominator := new(big.Int).Mul(reserveIn, big.NewInt(1000))
	denominator = new(big.Int).Add(denominator, amountInWithFee)
	amountOut := new(big.Int).Div(numerator, denominator)
	return amountOut
}
