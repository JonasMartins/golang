#!/bin/bash

# prod version with need of encrypt passwords

rm *.pem

# 1. generate ca private key
openssl req -x509 -newkey rsa:4096 -days 365 -keyout ca-key.pem -out ca-cert.pem -subj "/C=BR/ST=Ceara/L=Fortaleza/O=Backend/Dev/OU=Software/CN=/Admin/emailAddress=admin@gmail.com"

echo "CA's self signed certificate"
openssl x509 -in ca-cert.pem -noout -text

# 2. generate web server private key and signing request CSR
openssl req -newkey rsa:4096 -keyout server-key.pem -out server-cert.pem -subj "/C=BR/ST=Sao Paulo/L=Sao Paulo/O=Client/Engineering/OU=Hardware/CN=/Sam/emailAddress=sam@gmail.com"

# 3. use ca private key to sign web server csr and get back the signed cretificate
openssl x509 -req -in server-req.pem -days 60 -CA ca-cert.pem -CAKey ca-key.pem -CAcreateserial -out server-cert.pem -extfile server-ext.cnf

echo "Server's self signed certificate"
openssl x509 -in server-cert.pem -noout -text