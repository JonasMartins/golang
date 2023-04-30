#!/bin/bash

# dev version, without encrypt password

rm *.pem

# 1. generate ca private key
openssl req -x509 -newkey rsa:4096 -days 365 -nodes -keyout ca-key.pem -out ca-cert.pem -subj "/C=BR/ST=Sao Paulo/L=Sao Paulo/O=DEV/OU=TUTORIAL/CN=*.tutorial.dev/emailAddress=admin@gmail.com"

echo "CA's self signed certificate"
openssl x509 -in ca-cert.pem -noout -text

# 2. generate web server private key and signing request CSR
openssl req -newkey rsa:4096 -nodes -keyout server-key.pem -out server-req.pem -subj "/C=BR/ST=Sao Paulo/L=Sao Paulo/O=DEV/OU=BLOG/CN=*.clientservice.com/emailAddress=samuel@gmail.com"

# 3. use ca private key to sign web server csr and get back the signed cretificate
openssl x509 -req -in server-req.pem -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out server-cert.pem -extfile server-ext.conf


echo "Server's self signed certificate"
openssl x509 -in server-cert.pem -noout -text