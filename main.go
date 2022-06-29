package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	crp "github.com/libp2p/go-libp2p-core/crypto"
	peer "github.com/libp2p/go-libp2p-core/peer"
)

func main() {
	size := flag.Int("bitsize", 2048, "select the bitsize of the key to generate")
	typ := flag.String("type", "", "select type of key to generate (RSA, Ed25519, Secp256k1 or ECDSA)")
	key := flag.String("key", "", "specify the location of the key to decode it's peerID")

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

	var unmarshalPrivateKeyFucn func(data []byte) (crp.PrivKey, error)
	// rsa and ed25519 unmarshalPrivateKeyFucn are for backward compatibility
	// for keys saved with raw(), to read such keys, specify the key type
	switch strings.ToLower(*typ) {
	case "rsa":
		unmarshalPrivateKeyFucn = crp.UnmarshalRsaPrivateKey
	case "ed25519":
		unmarshalPrivateKeyFucn = crp.UnmarshalEd25519PrivateKey
	default:
		unmarshalPrivateKeyFucn = crp.UnmarshalPrivateKey
	}

	prvk, err := unmarshalPrivateKeyFucn(data)
	if err != nil {
		return err
	}

	id, err := peer.IDFromPrivateKey(prvk)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(os.Stderr, "Success!\nID for %s key: %s\n", prvk.Type().String(), id.Pretty())
	return err
}

func genKey(typ *string, size *int) error {
	var atyp int
	switch strings.ToLower(*typ) {
	case "rsa":
		atyp = crp.RSA
	case "ed25519":
		atyp = crp.Ed25519
	case "secp256k1":
		atyp = crp.Secp256k1
	case "ecdsa":
		atyp = crp.ECDSA
	default:
		return fmt.Errorf("unrecognized key type: %s", *typ)
	}

	fmt.Fprintf(os.Stderr, "Generating a %d bit %s key...\n", *size, *typ)

	priv, pub, err := crp.GenerateKeyPair(atyp, *size)
	if err != nil {
		return err
	}

	pid, err := peer.IDFromPublicKey(pub)
	if err != nil {
		return err
	}

	data, err := crp.MarshalPrivateKey(priv)
	if err != nil {
		return err
	}

	_, err = os.Stdout.Write(data)
	if err != nil {
		return nil
	}

	_, err = fmt.Fprintf(os.Stderr, "Success!\nID for generated key: %s\n", pid.Pretty())
	return err
}
