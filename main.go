package main

import (
	"flag"
	"fmt"
	"os"

	ci "github.com/ipfs/go-libp2p-crypto"
	peer "github.com/ipfs/go-libp2p-peer"
)

func main() {
	size := flag.Int("bitsize", 2048, "select the bitsize of the key to generate")
	typ := flag.String("type", "RSA", "select type of key to generate")

	flag.Parse()

	var atyp int
	switch *typ {
	case "RSA":
		atyp = ci.RSA
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
