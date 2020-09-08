package masterkey

import (
	"log"

	"git.roosoft.com/bitcoin/hd-wallets/lib"
	"github.com/tyler-smith/go-bip32"
)

var km *lib.KeyManager
var masterkey *bip32.Key

func init() {
	var err error

	km, err = lib.NewKeyManagerFromMnemonic(lib.Mnemonic, lib.Passphrase)
	if err != nil {
		log.Fatal(err)
	}

	masterkey, err = km.GetMasterKey()
	if err != nil {
		log.Fatal(err)
	}
}

// GetPrivateKey returns the private key
func GetPrivateKey() []byte {
	return masterkey.Key
}

// GetPublicKey returns the public key
func GetPublicKey() []byte {
	return masterkey.PublicKey().Key
}

// GetChainCode returns the extended private key's chain code
func GetChainCode() []byte {
	return masterkey.ChainCode
}
