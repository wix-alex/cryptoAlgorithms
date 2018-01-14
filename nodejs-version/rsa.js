var bignum = require('bignum');


const letters = ["a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q",
          "r","s","t","u","v","w","x","y","z",",",".","!","?",' '];
const numbers = [1,2,3,4,5,6,7,8,9,10,11,12,13,
          14,15,16,17,18,19,20,21,22,23,24,25,26,27,
          28,29,30,31];



function generateRSAKeyPair() {
    var p = bignum.prime(20);
    var q = bignum.prime(20);
    //if I put p q bigger, it takes too much time to decrypt
    p=15;
    q=17;
    var n = p * q;
    var phi = (p - 1) * (q - 1);
    var e = 65537;
    e=101;

    /*console.log("p=" + p + ", q=" + q);
    console.log("n=" + n);
    console.log("phi=" + phi);
    console.log("e=" + e);*/

    var pubK = {
        e: e,
        n: n
    };

    var d = modinv(e, phi);
    //console.log("d=" + d);

    var privK = {
        d: d,
        n: n
    };
    return({"pubK": pubK, "privK": privK});
}

function encrypt(m, pubK) {
    chars = m.split('');
    var r = [];
    for (var i=0; i<chars.length; i++) {
        numb = getNbyL(chars[i]);
        r.push(encryptNum(numb, pubK));
    }
    var r = encryptNum(numb, pubK)
    return r;
}
function decrypt(c, privK) {
    var r="";
    for (var i=0; i<c.length; i++) {
        numDecrypted = decryptNum(c[i], privK);
        r = r + getLbyN(numDecrypted);
    }
    return r;
}
function encryptNum(m, pubK) {
    var Me = bignum(m).pow(pubK.e);
    var c = bignum(Me).mod(pubK.n);
    return c;
}
function decryptNum(c, privK) {
    var Cd = bignum(c).pow(privK.d);
    var m = bignum(Cd).mod(privK.n);
    return m;
}
function getLbyN(n) {
    for (var i=0; i<numbers.length; i++) {
        if (numbers[i] == n) {
            return letters[i];
        }
    }
}
function getNbyL(l) {
    for (var i=0; i<letters.length; i++) {
        if (letters[i] == l) {
            return numbers[i];
        }
    }
}

// Extended Euclidean algorithm modified to get the Modular Multiplicative Inverse
function modinv(a, m) {
    var v = 1;
    var d = a;
    var u = (a == 1);
    var t = 1 - u;
    if (t == 1) {
        var c = m % a;
        u = Math.floor(m / a);
        while (c != 1 && t == 1) {
            var q = Math.floor(d / c);
            d = d % c;
            v = v + q * u;
            t = (d != 1);
            if (t == 1) {
                q = Math.floor(c / d);
                c = c % d;
                u = u + q * v;
            }
        }
        u = v * (1 - t) + t * (m - u);
    }
    return u;
}

module.exports = {
    generateKeys : generateRSAKeyPair,
    encrypt: encrypt,
    decrypt: decrypt
};
