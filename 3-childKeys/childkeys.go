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
func GetChild() (string, string, string) {
	compress := true // generate a compressed public key

	key, err := km.GetKey(lib.PurposeBIP49, lib.CoinTypeBTC, 0, 0, uint32(0))
	if err != nil {
		log.Fatal(err)
	}

	wif, _, segwitBech32, _, err := key.Encode(compress)
	if err != nil {
		log.Fatal(err)
	}

	return key.GetPath(), segwitBech32, wif
}
