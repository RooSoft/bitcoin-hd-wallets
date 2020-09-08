package seed

import (
	"log"

	"git.roosoft.com/bitcoin/hd-wallets/lib"
)

var km *lib.KeyManager

func init() {
	var err error

	km, err = lib.NewKeyManagerFromMnemonic(lib.Mnemonic, lib.Passphrase)
	if err != nil {
		log.Fatal(err)
	}
}

// GetSeed returns the seed
func GetSeed() []byte {
	seed := km.GetSeed()

	return seed
}

// GetMnemonic returns the mnemonic
func GetMnemonic() string {
	mnemonic := km.GetMnemonic()

	return mnemonic
}
