# hd-wallets

Bitcoin Wallet concetps PoC

- Create mnemonics
- Extract an xpub
- Extract a native segwit address

## How to run

```bash
go run main.go
```

## Example output

```
BIP39 Mnemonic:    trial utility cry style universe phone scorpion curious nasty zone rifle marriage random mercy pear salt skull burger helmet acid dentist tired tree ask
BIP39 Passphrase:  PASSWORD
BIP39 Seed:        d93097fc1fa5db6ea071e995105e932c2aed96e9f2da8e327c856894a03d90578d5af2e6957c0f9d498f16cc02004ee2aa5465742fe0389f8051d678253d1c90
BIP32 Root Key:    xprv9s21ZrQH143K4WKnR9kVTtR7xdszY89qRTzxNXeVNbxEgppW8XE2xbpopuS2ghV9iEb5jgpYwp2BPNzznJEt4bn8RVYKxD5N9ieaFbBDDiW
xpub:              xpub661MyMwAqRbcGzQFXBHVq2MrWfiUwasgngvZAv46vwVDZd9eg4YHWQ9HgC6xyTenAapgXv86ywjXaBbMwdeEQm5RL7tSBPQXTBGQvU4JM69
m/84'/0'/0'/0/0    bc1qr370xugpkxhyrqtxncdcmgc303efnm00fg4fu8 L4Wv2cUaGtoCFyLDBVHzd8YfXxERte6GaxqgzV9Ciy8LWCp3rCoc
```

## Inspiration

Based on [modood/btckeygen](https://github.com/modood/btckeygen)






seed
  f25bf8f4b0fb2c8f38768f8e419dcddfa5c3ce3bb4e9423869ed283aae4572a7
  83bc1e33f53dfb6424a765bd050b4b3ac4f8552babbab3376a12de987fd56b13

seed -> hmac   ---->   master private key
key        f8a7b2614be187565f87576f561f6870179ff8069bac0875a96f2e745d14cb2d
chaincode  11c5213cdfb6d89339c2c544a808569052c13957ae38c987a81de1a4c5e91480
