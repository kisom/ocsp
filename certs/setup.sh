#!/bin/bash

sign_cert () {
    echo "[+] signing cert $2 with profile $1"
    cfssl gencert -config config.json -ca ca.pem -ca-key ca-key.pem -profile $1 $2.json | \
	cfssljson -bare $2
}

cfssl gencert -initca ca.json | cfssljson -bare ca
sign_cert ocsp signer
sign_cert www  server
