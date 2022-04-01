pragma solidity ^0.8.12;

interface IUniswapV2Pair {
    function swap(uint amount0Out, uint amount1Out, address to, bytes calldata data) external;
    function token0() external view returns (address);
    function token1() external view returns (address);
}

interface IERC20 {

    function totalSupply() external view returns (uint256);
    function balanceOf(address account) external view returns (uint256);
    function allowance(address owner, address spender) external view returns (uint256);

    function transfer(address recipient, uint256 amount) external returns (bool);
    function approve(address spender, uint256 amount) external returns (bool);
    function transferFrom(address sender, address recipient, uint256 amount) external returns (bool);


    event Transfer(address indexed from, address indexed to, uint256 value);
    event Approval(address indexed owner, address indexed spender, uint256 value);
}


abstract contract UniswapV2Factory  {
    mapping(address => mapping(address => address)) public getPair;
    address[] public allPairs;
    function allPairsLength() external view virtual returns (uint);
}

contract F {
    address owner;

    constructor() {
        owner = msg.sender;
    }

    function uniswapV2Call(address sender, uint amount0, uint amount1, bytes calldata data) external {
        require(sender == owner, "Must be the owner.");
        require(amount0 == 0 || amount1 == 0, "Unidirectional trades only");

        address tokenToTransfer;
        // What do we need to pass in the calldata to know where to swap?
        // We pass all this into the calldata because we should be able to precalculate this of chain, saving gas costs.
        // 1. Address of the pair we should perform the swap on.
        // 2. The address of token0
        // 3. The address of token1
        // 4. The amount of tokens we expect back from the external pair
        // 5. The amount of tokens we should send back to the original pair to repay our flash swap.
        (address originalPair, address externalPair, address token0, address token1, uint amountOutExternalPair, uint amountToRepay) = abi.decode(data, (address, address, address, address, uint, uint));
        
        // Now we want to execute a swap on this external pair. Not a flash swap this time so we need to send the tokens to the pool first.
        // We send the token to the pool by calling the transfer function on the ERC20 token contract, in order to do this we will need the token address.
        {
            tokenToTransfer = amount0 == 0 ? token1 : token0;
            uint amountToSend = amount0 == 0 ? amount1 : amount0;
            uint amountOutExternal0 = amount0 == 0 ? 0 : amountOutExternalPair;
            uint amountOutExternal1 = amount1 == 0 ? 0 : amountOutExternalPair;
            require(IERC20(tokenToTransfer).transfer(externalPair, amountToSend), "Transfer of ERC20 failed to external pair");
            // Execute the trade to send us the amount of token we need.
            IUniswapV2Pair(externalPair).swap(amountOutExternal0, amountOutExternal1, address(this), "0x");
        }

        // Now we should own the amount of tokens that we need to repay our flash swap.
        // Send these tokens to the pair
        {
            tokenToTransfer = amount0 == 0 ? token0 : token1;
            require(IERC20(tokenToTransfer).transfer(originalPair, amountToRepay), "Repay ERC20 transfer");
        }
    }
}