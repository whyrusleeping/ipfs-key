# ipfs-key
A tool for easily generating ipfs keypairs. When run, it will write the bytes of
the serialized private key to stdout. By default, a 2048 bit RSA key will be
generated. The keysize can be changed by specifying the `-size` option, and the
key type can be changed by specifying the `-type` option (currently only RSA is
implemented).

## Installation
```
$ go get github.com/whyrusleeping/ipfs-key
```

## Usage
```
$ ipfs-key -size=4096 > my.key
```
