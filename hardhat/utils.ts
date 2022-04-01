import { BigNumber } from "ethers"

function getAmountIn(amountOut: any, reserveIn: any, reserveOut: any) {
    let numerator = reserveIn.mul(amountOut).mul(1000)
    let denominator = reserveOut.sub(amountOut).mul(997)
    return (numerator.div(denominator)) + 1
}

function getAmountOut(amountIn: BigNumber, reserveIn: BigNumber, reserveOut: BigNumber): BigNumber {
    let amountInWithFee = amountIn.mul(997)
    let numerator = amountInWithFee.mul(reserveOut)
    let denominator = reserveIn.mul(1000).add(amountInWithFee)
    return numerator.div(denominator)
}

module.exports = {
    getAmountIn,
    getAmountOut
}