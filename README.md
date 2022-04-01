# Trader Joe Arbitrage Bot

## Overview
Each block, AVAX/USDT reserves are queried from three DEXes (Trader Joe, Pangolin, Sushi Swap) and arbitrage opportunities are identified.

If an arbitrage opportunity is found then an optimal trade volume is determined by running a binary search algorithm.

If the trade is determined to be over the minimum acceptable profit, then the buy market pair's `swap` method is called with the calculated trade volume, a smart contract address, and calldata. 

Because of the passed calldata, the pair's `swap` method executes a flash swap by calling the `uniswapV2Call` (or `joeCall`/`pangolinCall`) method on the deployed smart contract `F.sol`

`F.sol` handles transfering the flash swapped tokens from the buy market to the sell market's pair, swapping for the correct number of output tokens, returning the necessary amount to the buy market in order to preserves the previous k value, and finally, transfering the profit to the EOA who initiated the transaction.

## Running the bot

First edit the assert statment at the top of the uniswapV2Call method in `F.sol` to allow for the correct DEXes to call this method.

Deploy `F.sol`

You will need a `.env` file containing your EOA's private key. A sample is provided at `.env.sample`

Then ensure you have Go installed with `go version`

[Download & Install Go](https://go.dev/doc/install)

Install dependencies with `go mod download`

Finally, start the bot with `go run main.go`

The `minimumProfit` variable can be changed in `main.go`

## Tests

Install dependencies with `npm install`

You will need a `.env` located at `hardhat/.env` containing your Alchemy key for an archive Ethereum node, a sample `.env` file is provided at `hardhat/.env.sample`

Tests can be ran from the `hardhat` directory with `npx hardhat test`
