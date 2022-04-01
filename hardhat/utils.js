"use strict";
exports.__esModule = true;
function getAmountIn(amountOut, reserveIn, reserveOut) {
    var numerator = reserveIn.mul(amountOut).mul(1000);
    var denominator = reserveOut.sub(amountOut).mul(997);
    return (numerator.div(denominator)) + 1;
}
function getAmountOut(amountIn, reserveIn, reserveOut) {
    var amountInWithFee = amountIn.mul(997);
    var numerator = amountInWithFee.mul(reserveOut);
    var denominator = reserveIn.mul(1000).add(amountInWithFee);
    return numerator.div(denominator);
}
module.exports = {
    getAmountIn: getAmountIn,
    getAmountOut: getAmountOut
};
