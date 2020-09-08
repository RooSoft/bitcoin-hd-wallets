package seed

import (
	"log"

	"git.roosoft.com/bitcoin/hd-wallets/lib"
)

var km *lib.KeyManager

func init() {
	pass := "PASSWORD"

	var err error

	km, err = lib.NewKeyManager(256, pass)
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
