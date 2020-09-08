package main

import (
	"fmt"
	"log"

	seed "git.roosoft.com/bitcoin/hd-wallets/1-seed"
	"git.roosoft.com/bitcoin/hd-wallets/lib"
)

// CoinType SLIP-0044 : Registered coin types for BIP-0044
// https://github.com/satoshilabs/slips/blob/master/slip-0044.md
type CoinType = uint32

const (
	CoinTypeBTC CoinType = 0x80000000
	CoinTypeLTC CoinType = 0x80000002
	CoinTypeETH CoinType = 0x8000003c
	CoinTypeEOS CoinType = 0x800000c2
)

type Purpose = uint32

const (
	PurposeBIP44 Purpose = 0x8000002C // 44' BIP44
	PurposeBIP49 Purpose = 0x80000031 // 49' BIP49
	PurposeBIP84 Purpose = 0x80000054 // 84' BIP84
)

func main() {
	pass := "PASSWORD"
	compress := true // generate a compressed public key

	km, err := lib.NewKeyManager(256, pass)
	if err != nil {
		log.Fatal(err)
	}

	masterKey, err := km.GetMasterKey()
	if err != nil {
		log.Fatal(err)
	}

	passphrase := km.GetPassphrase()
	if passphrase == "" {
		passphrase = "<none>"
	}

	fmt.Printf("%-18s %s\n", "BIP39 Mnemonic:", km.GetMnemonic())
	fmt.Printf("%-18s %s\n", "BIP39 Passphrase:", passphrase)
	fmt.Printf("%-18s %x\n", "BIP39 Seed:", km.GetSeed())
	fmt.Printf("%-18s %s\n", "BIP32 Root Key:", masterKey.B58Serialize())

	fmt.Printf("%-18s %x\n", "BIP32 key:", masterKey.Key)
	fmt.Printf("%-18s %x\n", "BIP32 chain code:", masterKey.ChainCode)

	publicKey := masterKey.PublicKey()
	fmt.Printf("%-18s %s\n", "xpub:", publicKey.String())

	key, err := km.GetKey(PurposeBIP84, CoinTypeBTC, 0, 0, uint32(0))
	if err != nil {
		log.Fatal(err)
	}
	wif, _, segwitBech32, _, err := key.Encode(compress)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%-18s %s %s\n", key.GetPath(), segwitBech32, wif)

	/// see https://learnmeabitcoin.com/technical/hd-wallets

	fmt.Printf("\n\n1. Seed\n")
	fmt.Printf("-------\n")
	fmt.Printf("%-10s %x\n", "seed", seed.GetSeed())
	fmt.Printf("%-10s %s\n", "mnemonic", seed.GetMnemonic())

}
