//SPDX-License-Identifier: Unlicense
pragma solidity ^0.8.11;

import "hardhat/console.sol";

interface IUniswapV2Pair {
    function swap(uint amount0Out, uint amount1Out, address to, bytes calldata data) external;
    function token0() external view returns (address);
    function token1() external view returns (address);
}

interface IUniswapV2Factory {
    event PairCreated(address indexed token0, address indexed token1, address pair, uint);
    function feeTo() external view returns (address);
    function feeToSetter() external view returns (address);
    function getPair(address tokenA, address tokenB) external view returns (address pair);
    function allPairs(uint) external view returns (address pair);
    function allPairsLength() external view returns (uint);
    function createPair(address tokenA, address tokenB) external returns (address pair);
    function setFeeTo(address) external;
    function setFeeToSetter(address) external;
}

interface IERC20 {
    function totalSupply() external view returns (uint256);
    function balanceOf(address account) external view returns (uint256);
    function allowance(address owner, address spender) external view returns (uint256);
    function transfer(address recipient, uint256 amount) external returns (bool);
    function approve(address spender, uint256 amount) external returns (bool);
    function transferFrom(address sender, address recipient, uint256 amount) external returns (bool);
    function decimals() external returns (uint256);
    event Transfer(address indexed from, address indexed to, uint256 value);
    event Approval(address indexed owner, address indexed spender, uint256 value);
}

interface SPECIALTETHERIERC20 {
    function totalSupply() external view returns (uint256);
    function balanceOf(address account) external view returns (uint256);
    function allowance(address owner, address spender) external view returns (uint256);
    function transfer(address recipient, uint256 amount) external;
    function approve(address spender, uint256 amount) external returns (bool);
    function transferFrom(address sender, address recipient, uint256 amount) external returns (bool);
    event Transfer(address indexed from, address indexed to, uint256 value);
    event Approval(address indexed owner, address indexed spender, uint256 value);
}

contract F {
    address owner;
    address EuniFactory = 0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f;
    address EsushiFactory = 0xC0AEe478e3658e2610c5F7A4A2E1777cE9e4f2Ac;
    address joeFactory = 0x9Ad6C38BE94206cA50bb0d90783181662f0Cfa10;
    address pangolinFactory = 0xefa94DE7a4656D787667C749f7E1223D71E9FD88;
    address sushiFactory = 0xc35DADB65012eC5796536bD9864eD8773aBc74C4;

    constructor() {
        owner = msg.sender;
    }
    
    function uniswapV2Call(address sender, uint amount0, uint amount1, bytes calldata data) external {
        require(sender == owner, "Only owner may call.");
        address token0 = IUniswapV2Pair(msg.sender).token0(); // fetch the address of token0
        address token1 = IUniswapV2Pair(msg.sender).token1(); // fetch the address of token1

        // Need to add more OR statements here if adding more exchanges.
        assert(
            // msg.sender == IUniswapV2Factory(joeFactory).getPair(token0, token1) ||
            // msg.sender == IUniswapV2Factory(pangolinFactory).getPair(token0, token1) ||
            //msg.sender == IUniswapV2Factory(sushiFactory).getPair(token0, token1)
            msg.sender == IUniswapV2Factory(EsushiFactory).getPair(token0, token1) ||
            msg.sender == IUniswapV2Factory(EuniFactory).getPair(token0, token1)
        );

        (address originalPair, address externalPair, uint amountOutExternalPair, uint amountToRepay) = abi.decode(data, (address, address, uint256, uint256));
        {
            SPECIALTETHERIERC20 tToTransfer = SPECIALTETHERIERC20(amount0 == 0 ? token1 : token0);
            tToTransfer.transfer(externalPair, amount0 == 0 ? amount1 : amount0);
            IUniswapV2Pair(externalPair).swap(amount0 == 0 ? amountOutExternalPair : 0, amount1 == 0 ? amountOutExternalPair : 0, address(this), "");
        }
        {
            IERC20 t = IERC20(amount0 == 0 ? token0 : token1);
            t.transfer(originalPair, amountToRepay);
            uint num = t.balanceOf(sender);
            t.transfer(sender, t.balanceOf(address(this)));
            require(ensureProfit(num, t, sender), "Trade was not profitable");
        }
    }

    function ensureProfit(uint initialBalance, IERC20 token, address sender) internal view returns(bool) {
        uint finalBalance = token.balanceOf(sender);
        return finalBalance > initialBalance;
    }

    // *** Basically just aliases for what each DEX calls their callback function. ***

    function joeCall(address sender, uint256 amount0, uint256 amount1, bytes calldata data) external {
        this.uniswapV2Call(sender, amount0, amount1, data);
    }
    
    function pangolinCall(address sender, uint256 amount0, uint256 amount1, bytes calldata data) external {
        this.uniswapV2Call(sender, amount0, amount1, data);
    }
}
