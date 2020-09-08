package childkeys

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

// GetChild returns the seed
func GetChild(purpose, coinType, account, change, index uint32) (string, string, string, string, string) {
	compress := true // generate a compressed public key

	key, err := km.GetKey(purpose, coinType, account, change, index)
	if err != nil {
		log.Fatal(err)
	}

	wif, address, segwitBech32, segwitNested, err := key.Encode(compress)
	if err != nil {
		log.Fatal(err)
	}

	return key.GetPath(), address, segwitBech32, segwitNested, wif
}
