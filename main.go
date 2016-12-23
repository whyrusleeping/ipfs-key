package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	peer "gx/ipfs/QmfMmLGoKzCHDN7cGgk64PJr4iipzidDRME8HABSJqvmhC/go-libp2p-peer"
	ci "gx/ipfs/QmfWDLQjGjVe4fr5CoztYW2DYYjRysMJrFe1RCsXLPTf46/go-libp2p-crypto"
)

func main() {
	size := flag.Int("bitsize", 2048, "select the bitsize of the key to generate")
	typ := flag.String("type", "RSA", "select type of key to generate (RSA or Ed25519)")

	flag.Parse()

	var atyp int
	switch strings.ToLower(*typ) {
	case "rsa":
		atyp = ci.RSA
	case "ed25519":
		atyp = ci.Ed25519
	default:
		fmt.Fprintln(os.Stderr, "unrecognized key type: ", *typ)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stderr, "Generating a %d bit %s key...\n", *size, *typ)
	priv, pub, err := ci.GenerateKeyPair(atyp, *size)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Fprintln(os.Stderr, "Success!")

	pid, err := peer.IDFromPublicKey(pub)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stderr, "ID for generated key: %s\n", pid.Pretty())

	data, err := priv.Bytes()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	os.Stdout.Write(data)
}
