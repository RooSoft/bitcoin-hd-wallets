package seed

import (
	"log"

	"git.roosoft.com/bitcoin/hd-wallets/lib"
	"github.com/tyler-smith/go-bip32"
)

var km *lib.KeyManager
var masterKey *bip32.Key

func init() {
	var err error

	km, err = lib.NewKeyManagerFromMnemonic(lib.Mnemonic, lib.Passphrase)
	if err != nil {
		log.Fatal(err)
	}

	masterKey, err = km.GetMasterKey()
	if err != nil {
		log.Fatal(err)
	}
}

// GetXpriv returns the xpriv
func GetXpriv(purpose, coinType, account, change, index uint32) string {
	xpriv := masterKey.B58Serialize()

	return xpriv
}

// GetXpub returns the xpub
func GetXpub(purpose, coinType, account, change, index uint32) string {
	publicKey := masterKey.PublicKey()

	return publicKey.B58Serialize()
}
