require("@nomiclabs/hardhat-waffle");
const dotenv = require('dotenv')
dotenv.config()

// "https://eth-mainnet.alchemyapi.io/v2/" + process.env.ALCHEMY_KEY
// "https://0jrmwv1jvhx4.usemoralis.com:2053/server"
module.exports = {
  solidity: "0.8.11",
  networks: {
    hardhat: {
      forking: {
        url: "https://eth-mainnet.alchemyapi.io/v2/" + process.env.ALCHEMY_KEY,
        blockNumber: 14219588
      }
    }
  }
};
