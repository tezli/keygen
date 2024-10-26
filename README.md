# Keygen
![build](https://github.com/tezli/keygen/actions/workflows/main.yml/badge.svg)
[![codecov](https://codecov.io/gh/tezli/keygen/branch/main/graph/badge.svg)](https://codecov.io/gh/tezli/keygen)

Simple tool for genrating crytographic keys.

# Installation

**For the gophers**
```shell
go install github.com/tezli/keygen
```
**For eveyone else**

```shell
curl -so "https://github.com/tezli/keygen/releases/download/v0.0.1/keygen-$(uname -s | tr A-Z a-z)-$(uname -m | sed s/x86_/amd/g)" /usr/local/bin/keygen && chmod +x /usr/local/bin/keygen
```

## Usage

**Creating an ECDSA key**
```shell
keygen ecdsa --curve P521 --out ecdsa_key
```
or short:
```shell
keygen ecdsa -c P521 -o ecdsa_key
```
Add the ` -p ` flag to create a public key also.

**Creating a RSA key**
```shell
keygen rsa --bits 4096 --out rsa_key
```
or short:
```shell
keygen ecdsa -b 4096 -o rsa_key
```

**Creating key with a corresponding public key**
```shell
keygen rsa -p -b 4096 -o rsa_key
```
will generate:
```shell
rsa_key
rsa_key.pub
```
 
