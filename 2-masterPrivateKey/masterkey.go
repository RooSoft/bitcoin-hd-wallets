package masterkey

import (
	"log"

	"git.roosoft.com/bitcoin/hd-wallets/lib"
	"github.com/tyler-smith/go-bip32"
)

var km *lib.KeyManager
var masterkey *bip32.Key

func init() {
	pass := "PASSWORD"

	var err error

	km, err = lib.NewKeyManager(256, pass)
	if err != nil {
		log.Fatal(err)
	}

	masterkey, err = km.GetMasterKey()
	if err != nil {
		log.Fatal(err)
	}
}

// GetKey returns the private key
func GetKey() []byte {
	return masterkey.Key
}

// GetChainCode returns the extended private key's chain code
func GetChainCode() []byte {
	return masterkey.ChainCode
}
