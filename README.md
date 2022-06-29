# ipfs-key

[![standard-readme compliant](https://img.shields.io/badge/readme%20style-standard-brightgreen.svg?style=flat-square)](https://github.com/RichardLitt/standard-readme)

> A tool for easily generating and reading ipfs keypairs

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Contribute](#contribute)
- [License](#license)

## Installation

```
$ go get github.com/whyrusleeping/ipfs-key
```

## Usage

When run, it will write the bytes of
the serialized private key to stdout. By default, a 2048 bit RSA key will be
generated. In this case the key size can be changed by specifying the `-bitsize`
option. The key type can be changed by specifying the `-type` option (RSA, Ed25519, Secp256k1 or ECDSA).

```
$ ipfs-key -bitsize=4096 > my-rsa4096.key
Generating a 4096 bit RSA key...
Success!
ID for generated key: QmS5cwbxmGyPiEH3SYNgiAazG46NvogKxfx2iX6jt4ef1a
```
```
$ ipfs-key -type=ed25519 > my-ed.key
Generating a 2048 bit ed25519 key...
Success!
ID for generated key: 12D3KooWHM4kLNwS2FzN5GtG5Dfy9h7dLTRs3rtuF9NiR4mjBv3h
```
```
$ ipfs-key -key my-ed.key
Reading key at: my-ed.key
Success!
ID for ed25519 key: 12D3KooWF1TKgiqLMh14za7dWMN5RFRC1WAvgHYioksmdwuhZkzT
```
For backward compatibility, to read RSA and Ed25519 keys generated with raw(), specify the `-type rsa` or `-type ed25519` before the `-key` 
```
$ ipfs-key --type rsa -key my-ed.key
Reading key at: my-ed.key
Success!
ID for rsa key: 12D3KooWF1TKgiqLMh14za7dWMN5RFRC1WAvgHYioksmdwuhZkzT
```

## Contribute

PRs accepted.

## License

[MIT](LICENSE) Copyright (c) 2016 [Jeromy Johnson](http://github.com/whyrusleeping)
