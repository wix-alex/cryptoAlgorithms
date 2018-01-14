var rsa = require('./rsa');

console.log("RSA test");

var rsaKeyPair = rsa.generateKeys();
console.log(rsaKeyPair);

var m = "hi, trying the rsa encryption";
console.log("m=" + m);

var c = rsa.encrypt(m, rsaKeyPair.pubK);
console.log("c encrypted: " + c);

var m2 = rsa.decrypt(c, rsaKeyPair.privK);
console.log("m decrypted: " + m2);



console.log("-----");
console.log("Paillier test");


var paillier = require('./paillier');

var paillierKeyPair = paillier.generateKeys();
console.log(paillierKeyPair);

var m = "hi, trying the paillier encryption";
console.log("m=" + m);

var c = paillier.encrypt(m, paillierKeyPair.pubK);
console.log("c encrypted: " + c);

var m2 = paillier.decrypt(c, paillierKeyPair.pubK, paillierKeyPair.privK);
console.log("m decrypted: " + m2);
