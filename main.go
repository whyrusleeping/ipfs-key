package main

import (
	// Update the import path for crypto
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	crp "github.com/libp2p/go-libp2p/core/crypto"
	peer "github.com/libp2p/go-libp2p/core/peer"
)

func main() {
	size := flag.Int("bitsize", 2048, "select the bitsize of the key to generate")
	typ := flag.String("type", "", "select the type of key to generate (RSA, Ed25519, Secp256k1, or ECDSA)")
	key := flag.String("key", "", "specify the location of the key to decode its peer ID")

	flag.Parse()

	if *key != "" {
		if err := readKey(key, typ); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		return
	}

	if *typ == "" {
		*typ = "RSA"
	}
	if err := genKey(typ, size); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	return
}

func readKey(keyLoc *string, typ *string) error {
	data, err := ioutil.ReadFile(*keyLoc)
	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Reading key at: %s\n", *keyLoc)

	var unmarshalPrivateKeyFunc func(data []byte) (crp.PrivKey, error)
	switch strings.ToLower(*typ) {
	case "rsa":
		unmarshalPrivateKeyFunc = crp.UnmarshalRsaPrivateKey
	case "ed25519":
		unmarshalPrivateKeyFunc = crp.UnmarshalEd25519PrivateKey
	default:
		unmarshalPrivateKeyFunc = crp.UnmarshalPrivateKey
	}

	prvKey, err := unmarshalPrivateKeyFunc(data)
	if err != nil {
		return err
	}

	id, err := peer.IDFromPrivateKey(prvKey)
	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Success!\nID for %s key: %s\n", prvKey.Type().String(), id)
	return nil
}

func genKey(typ *string, size *int) error {
	var keyType int
	switch strings.ToLower(*typ) {
	case "rsa":
		keyType = crp.RSA
	case "ed25519":
		keyType = crp.Ed25519
	case "secp256k1":
		keyType = crp.Secp256k1
	case "ecdsa":
		keyType = crp.ECDSA
	default:
		return fmt.Errorf("unrecognized key type: %s", *typ)
	}

	fmt.Fprintf(os.Stderr, "Generating a %d-bit %s key...\n", *size, *typ)

	prvKey, pubKey, err := crp.GenerateKeyPair(keyType, *size)
	if err != nil {
		return err
	}

	id, err := peer.IDFromPublicKey(pubKey)
	if err != nil {
		return err
	}

	data, err := crp.MarshalPrivateKey(prvKey)
	if err != nil {
		return err
	}

	_, err = os.Stdout.Write(data)
	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Success!\nID for the generated key: %s\n", id)
	return nil
}
