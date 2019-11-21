#!/usr/bin/env bash

usage() {
    echo "usage: gen.sh <your domain>"
}

gen() {
    domain=$1
    echo "--------------------------------"
    echo ">>> start to gen certs for ${domain}"
    echo "--------------------------------"
    echo ">>> generate key ${domain}.key"
    openssl genrsa -out ${domain}.key 2048
    echo "--------------------------------"
    echo ">>> request  csr ${domain}.csr"
    openssl req -new -key ${domain}.key -out ${domain}.csr -subj "/CN=${domain}"
    echo "--------------------------------"
    echo ">>> generate crt ${domain}.crt"
    openssl x509 -req -days 365 -in ${domain}.csr -signkey ${domain}.key -out ${domain}.crt
    echo "--------------------------------"
    echo ">>> generation done. file list:"
    echo "  ${domain}.key"
    echo "  ${domain}.csr"
    echo "  ${domain}.crt"
    echo ""
    echo ">>> To create secret tls, use command:"
    echo "  kubectl create secret tls --cert ${domain}.crt --key ${domain}.key <your_tls_name> -n <namespace>"
    echo "--------------------------------"
}

domain=$1

if [[ ! -n ${domain} ]] ;
then
    usage
else
    gen ${domain}
fi