var bignum = require('bignum');

const keyLength = 10;

var keyPair = generatePaillierKeyPair();
var pubK = keyPair.pubK;
var privK = keyPair.privK;
console.log(keyPair);

function generatePaillierKeyPair(){
    var p = bignum.prime(keyLength/2);
    var q = bignum.prime(keyLength/2);
    //if I put p q bigger, it takes too much time to decrypt
    p=15;
    q=17;
    var compr = gcd(p*q, (p-1)*(q-1));
    console.log("gcd(p*q, (p-1)*(q-1)) = " + compr);
    console.log("p: " + p + ", q: " + q);

    var n = p * q;
    console.log("n= " + n);

    var lambda = lcm(p-1, q-1);
    console.log("lambda= " + lambda);

    var alpha = rand(0, n);
    var beta = rand(0, n);
    console.log("alpha - beta: " + alpha + " - " + beta);
    var g = bignum(bignum(alpha * n + 1).mul(bignum(beta).pow(n)).mod(n*n));
    console.log("g: " + g);

    var pubK = {
        n: n,
        g: g.toNumber()
    };
    console.log("pubK:");
    console.log(pubK);

    console.log(bignum(1).div(bignum(L(bignum(g.pow(lambda)).mod(n*n).toNumber(), n))));
    var mu = L(bignum(g.pow(lambda)).mod(n*n).toNumber(), n).pow(-1).mod(n);
    console.log("mu: " + mu);


    var privK = {
        lambda: lambda,
        mu: mu.toNumber()
    };
    console.log("privK:");
    console.log(privK);
    return({"pubK": pubK, "privK": privK});
}

function rand(min, max) {
    return Math.floor(Math.random() * (max - min + 1)) + min;
}
function L(u, n){
    return(bignum((u-1)/n));
}
function gcd(a, b) {
    if ( ! b) {
        return a;
    }

    return gcd(b, a % b);
}

function lcm(a, b) {
    return( (a*b) / gcd(a, b));
}
function random(bitLength) {

      var wordLength = bitLength / 4 / 8;

      var randomWords = sjcl.random.randomWords(wordLength),
          randomHex = sjcl.codec.hex.fromBits(randomWords);

      return new BigInteger(randomHex, 16);

}
