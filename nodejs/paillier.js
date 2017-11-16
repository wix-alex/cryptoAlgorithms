var bignum = require('bignum');

const letters = ["a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q",
    "r", "s", "t", "u", "v", "w", "x", "y", "z", ",", ".", "!", "?", ' '
];
const numbers = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13,
    14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27,
    28, 29, 30, 31
];
const keyLength = 10;



function generatePaillierKeyPair() {
    var p = bignum.prime(keyLength / 2);
    var q = bignum.prime(keyLength / 2);
    //if I put p q bigger, it takes too much time to decrypt
    p = 29;
    q = 479;
    var compr = gcd(p * q, (p - 1) * (q - 1));
    /*console.log("gcd(p*q, (p-1)*(q-1)) = " + compr);
    console.log("p: " + p + ", q: " + q);*/

    var n = bignum(p * q);
    //console.log("n= " + n);

    var lambda = lcm(p - 1, q - 1);

    var alpha = bignum(rand(0, n.toNumber()));
    var beta = bignum(rand(0, n.toNumber()));
    var alphan = alpha.mul(n);
    var alphan1 = alpha.add(1);
    var betaN = beta.pow(n);
    var ab = alphan1.mul(betaN);
    var n2 = n.pow(2);
    var g = ab.mod(n2);
    //in some Paillier implementations use this:
    //var g = n.add(1);

    var pubK = {
        n: n,
        g: g
    };

    //var lFunc = L(bignum(g.pow(lambda)).mod(n*n).toNumber(), n);
    var Glambda = g.pow(lambda);
    var u = Glambda.mod(n2);
    var lFunc = L(u, n);
    var mu = lFunc.invertm(n);


    var privK = {
        lambda: bignum(lambda),
        mu: mu
    };
    return ({
        "pubK": pubK,
        "privK": privK
    });
}



function encrypt(m, pubK) {
    chars = m.split('');
    var r = [];
    for (var i = 0; i < chars.length; i++) {
        numb = getNbyL(chars[i]);
        r.push(encryptNum(numb, pubK));
    }
    return r;
}

function decrypt(c, pubK, privK) {
    var r = "";
    for (var i = 0; i < c.length; i++) {
        numDecrypted = decryptNum(c[i], pubK, privK);
        r = r + getLbyN(numDecrypted);
    }
    return r;
}

function encryptNum(m, pubK) {
    m = bignum(m);
    gM = pubK.g.pow(m);
    r = bignum(rand(0, pubK.n.toNumber()));
    rN = r.pow(pubK.n);
    gMrN = gM.mul(rN);
    n2= pubK.n.pow(2);
    c = gMrN.mod(n2);
    return c.toNumber();
}

function decryptNum(c, pubK, privK) {
    c = bignum(c);
    cLambda = c.pow(privK.lambda);
    n2 = pubK.n.pow(2);
    u = cLambda.mod(n2);
    lFunc = L(u, pubK.n);
    LMu = lFunc.mul(privK.mu);
    m = LMu.mod(pubK.n);
    return m;
}

function getLbyN(n) {
    for (var i = 0; i < numbers.length; i++) {
        if (numbers[i] == n) {
            return letters[i];
        }
    }
}

function getNbyL(l) {
    for (var i = 0; i < letters.length; i++) {
        if (letters[i] == l) {
            return numbers[i];
        }
    }
}

function rand(min, max) {
    return Math.floor(Math.random() * (max - min + 1)) + min;
}

function L(u, n) {
    u1 = u.sub(bignum(1));
    l = u1.div(n);
    return (l);
}

function gcd(a, b) {
    if (!b) {
        return a;
    }

    return gcd(b, a % b);
}

function lcm(a, b) {
    return ((a * b) / gcd(a, b));
}

function random(bitLength) {

    var wordLength = bitLength / 4 / 8;

    var randomWords = sjcl.random.randomWords(wordLength),
        randomHex = sjcl.codec.hex.fromBits(randomWords);

    return new BigInteger(randomHex, 16);

}


module.exports = {
    generateKeys : generatePaillierKeyPair,
    encrypt: encrypt,
    decrypt: decrypt
};
