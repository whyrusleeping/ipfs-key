# ipfs-key

[![standard-readme compliant](https://img.shields.io/badge/readme%20style-standard-brightgreen.svg?style=flat-square)](https://github.com/RichardLitt/standard-readme)

> A tool for easily generating ipfs keypairs

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
option. The key type can be changed by specifying the `-type` option (rsa or
ed25519).

```
$ ipfs-key -bitsize=4096 > my-rsa4096.key
$ ipfs-key -type=ed25519 > my-ed.key
```

## Contribute

PRs accepted.

## License

[MIT](LICENSE) Copyright (c) 2016 [Jeromy Johnson](http://github.com/whyrusleeping)
