# hd-wallets

Bitcoin Wallet concetps PoC

1- Create mnemonics and seed
2- Create a master key with chain code
3- Derive some child keys
4- Extract an xpub and xpriv

## How to run

```bash
go run main.go
```

## Example output

```
1. Seed
-------
mnemonic  glow laugh acquire menu anchor evil occur put hover renew calm purpose
seed      b1680c7a6ea6ed5ac9bf3bc3b43869a4c77098e60195bae51a94159333820e125c3409b8c8d74b4489f28ce71b06799b1126c1d9620767c2dadf642cf787cf36


2. Master key
-------------
private key  081549973bafbba825b31bcc402a3c4ed8e3185c2f3a31c75e55f423e9629aa3
chain code   1d7d2a4c940be028b945302ad79dd2ce2afe5ed55e1a2937a5af57f8401e73dd
public key   0343b337dec65a47b3362c9620a6e6ff39a1ddfa908abab1666c8a30a3f8a7cccc


3. Child keys
-------------

---BIP 44 ---
m/44'/0'/0'/0/0  1Mx2LZjB78QF46bYoYUG9ok62DHPTnoyhv 3EHpg2DVseJQ2ghekfTav1reifzz5gG722 bc1quh9g9z2h5gtt68y07lkjzls0pvf59gfuragjlj
m/44'/0'/0'/0/1  1E3EiavU8L9G8EQ8dxscaRYMW5vVHhDM4e 3Lx6uT2LBD8jt44osmBpDAJQ6ZWgwvtZsE bc1q3uz74hm9yrs7x7gx9ssd3f8x8vpq7u9gwf0c63
m/44'/0'/0'/0/2  13CtFFhoH22uoLHCBKpew8xcsUE6BacsmP 3MDqESJaq9UtmRXHij57zRq5qegDsvXtKS bc1qrqhe4hd0k9uaau83aytywktnqf0ayw3zw9vaen
m/44'/0'/0'/0/3  1MeSPRbdx4LEvQvuL7LRrwafvLFFKLB9bz 3A8QyRq986wcRp5ofFbBFF925a1fttyWda bc1qufm3jyc2wy5xf83qhf4hkmvcm98w7g0jqfd0uz
m/44'/0'/0'/0/4  1Fjd76Sjmss93EpGvirEJ4H9cn8oZSQxcc 39mpYFmcWN7UETWvf1HeErJ9EsmNz1JuNk bc1q5xs4anl9ujd8havhuz57p4rqdec23mqdp0r8rg
m/44'/0'/0'/0/5  17sUD8YFZZ3nnFXgrpyheuzyUdLa7Hpywj 3HqxWSUNTczRqsnP1GiSsPR5dqKCiZqxrE bc1qfdwygrw6qp8pn222md3cwc3rzxjkj82akcjy0n
m/44'/0'/0'/0/6  1Q1fBeK7KU6nkCgEEcu2pKDV7CppDd9hQA 32JsGB1vrJW4pk3D397LsDDvmSp6zpvZR5 bc1ql34vgkhxhl7n83sjksvz9weym8dca6qwk27hp2
m/44'/0'/0'/0/7  1Hg4ussFTdefZvkzn61NT8ShGh6se8LKS2 3CD3HLhFW193CKKvyDcKXvf5b2GrkKPG1u bc1qkmjcq9hsyy70h3ywps74h3es8sxqcdh9k4efkw
m/44'/0'/0'/0/8  1BfWLL8M1uqt7cRBT87Hwqbjxvz6kojNpK 3KD7UUUdgKw5gDoNyyCSZBWQW18kVSf71u bc1qwnuceusxk66dzdw3kzp97tpv37fh0z325yxcv6
m/44'/0'/0'/0/9  17VJwra6VjQkBhgqU3gCZo7KN9wfTYdftG 336EeBxMbfqy8EBNqwmTRWubYzfDkxQDFm bc1qgu4krjeeuyh42se6ynhal6rf3gj46uswu0xdua
m/44'/0'/0'/0/10 14aMMGf2azQYxrihH7ny8e45BLNYXkXj3t 35bQfqBSh4Zzem3BqnJUKRKgyutotsuWT8 bc1qyum3wj5esjdu29vrpjpx6p2erjwqd2vezvkjjq
m/44'/0'/0'/0/11 1HNKSQUz4dtH99KRpsPsiCLFCqBzDwGKcf 3AqJt8HaQgd4uV9jXXbdMrWiUhtEZmDpm4 bc1qkw9z894vvh7nd6cx6cwvtnjle3nydx5dsz3wka
m/44'/0'/0'/0/12 18WPDFyT5w2w6ZY2uBm9ppYfS9b2ZRQNd7 38mZJ11ZdtejG9KJxV16GSDAcPYu4sfU21 bc1q2fthcp4cmej2y0ludpx8snaqjv7jj54t3tzwv6
m/44'/0'/0'/0/13 1JSThgDDDEjRbY9aPxLSqsSTQDY3UzTF5S 3AWgZDuh5CdiffY19fsMi6avHjdVK9nGeg bc1qha9d89v30jwcxyrewphuvtszw4pg7zva4pcqar
m/44'/0'/0'/0/14 1LmYrocvq6HKsLqtrmwNhZiGwAgYE2uoVP 3712yFLpvY9Cr6KXsVae86cJGK8K4wvAkb bc1qmrt5jkp4m89ahqx526695wah7r35jmklk03h33
m/44'/0'/0'/0/15 16GzXPUsVT4R3ws81S2dFh1LqhGB2ggpJs 38eBeZTGqFucVLQLahcfJdxNx1a1MF9GmK bc1q880wxz744fl0amhx030jky48dh0ejhpvccx2dn
m/44'/0'/0'/0/16 16T8N9S3PtZXQJn4y2PqdbsdHtZnFhZQCY 3DGS7jmTcS3a9w2NV9ZEkEwB2tthS6WVVk bc1q80yewu9jkz4d4fw2aakmzsauahdc5dcxjdszx0
m/44'/0'/0'/0/17 1NqiX8DsBQkBNk1xNzWJPmbS13Dj17UaWT 3BweVH4CdRR98Zi7Fgg3cp397GeFsyZNZH bc1qa7g5g20jsr3nt2cm05x2gvqaeq9xnvlr54qr3t
m/44'/0'/0'/0/18 1GnDwxjFR8ZJK6uHX8siGMdNkPnZgVjMR6 34NJLL4SJazJ5VXakaPzd6PawJWTEyvN5y bc1q45tkn3psnayn45mnzwjpkafs64ygqqajpcyaq7
m/44'/0'/0'/0/19 1M6s38mo47E6uXXgZ85AfUBE4am527KqQ4 37BTwvVaTV2iYvHZinpQLU9y2weixXaLEq bc1qm3l96dt92myg7f2p2wrnd4heq2e6lzv9j0t0c3


---BIP 84 ---
m/84'/0'/0'/0/0  1BvybryayE2PUdrFrsrxrrLvwtPRAd5z9  3DghCyw3JqzdAA9YUSYNE2PrSofivs2aaP bc1qqggnweh9rgafgzc7kv64dk6g6up5704szk3j5x
m/84'/0'/0'/0/1  154HGkys4aQYYStGq3xdVE5oi9S1L61sFk 3BgkAyE2CLpNhummbEpr7rBiztqDaTiq6d bc1q93l42lvd5x30xn0s94z0jpvtuw8rzln8ap7da6
m/84'/0'/0'/0/2  1N8AztxXgfuCVn1VkMgJr982B8J13U4JqZ 394KYjmEv2ZQ6nRdXCVBFjVuQ5RnJ3W8ef bc1qu767vd439hfcjvtsrdq9wrlm7d9crasfpl9sc2
m/84'/0'/0'/0/3  1EJnckvCWcrHznADkVLgLiTpV7YiZRyEyt 39VFhSYtTKrNrEx5XKPZ9oJDCsySM5PaLj bc1qj8mvjzqftf3sc6d9whdpgvwvuudtrzen4klfxd
m/84'/0'/0'/0/4  1A9PCWCajEVBtss9bKL4UVujwdmAwGMaeP 39kpZSKuFeKz56VYyjohasZmVjiWdkSDpC bc1qv3806maw3qurrhnqqa9wnf3vuvl5ps9gnrdyjq
m/84'/0'/0'/0/5  1KXSqpFvzwWJ6wejxbVLf7dRE7tseGaRFy 3EHQEoh3KKC21oCyQ9g2C2jy7rDGtj6hCH bc1qev69pqf4dxwgpk0n0mxstcya7qhfv7s8ygusd4
m/84'/0'/0'/0/6  1NibcHaWG3gJwvf2n1ebc247DHb15AUB65 394u14DJzVxuDg4kpPG1TxBYV7w65vj9ow bc1qacufteqqe7k3huzj4hcwjwfkyjzkxv28x2t7kz
m/84'/0'/0'/0/7  1MZU3sJMrhAvBmP24ThuPzaQubyTGWhygU 3Njy84SUdjns9Xe3ysAnyBmXTqu4jZCXjh bc1quxrx07adx77knzpch9l2c3ys0ts5jq0pq92tjm
m/84'/0'/0'/0/8  122pjReexKWXQJYYUPKwMC6isNEfSHrATr 325L4D6CvqNxtaa7B1SGigbGS51JEZVa9W bc1qpdgx9n40phncpyn5vxcuf77uy6y23ed9wm4hvn
m/84'/0'/0'/0/9  178xWJcRxgLSpKDPCiAfyNdTTJRXQt8HL3 39XJcdkLP9KMbRkW1rVqXJjTQ6cdEMSFEz bc1qgdgllau7dlssv99jfycjmxcypzd6tegrrusfaf
m/84'/0'/0'/0/10 186LqitsdPuNakYLxGbudbBB6GZ4gFEqYq 3Qi3xP1qy4JxJwF7zVgnYRRKmcNDuV4TH3 bc1qfh9ct88uuf5e766e9llpwd0j2549d2jc3wjksc
m/84'/0'/0'/0/11 1M29xHdsyqETq8rvndAeLqCpX6ziUyHfcq 39JDLEiuVPksbFQdwH5fFnP5Dj6U3N49PT bc1qmwdx0j8jvvv3tp642gkwlrtlytj0uullu82g4c
m/84'/0'/0'/0/12 1PRKdce7NjvhRbmJVasy5msT4c4DVrf9QA 3KgdzvBiNN5j8vs8bHFjqLtfxJ1yuzEqRG bc1q7hkymrtxppwpz25h68ts2gy2x5c5cfdynq0t9y
m/84'/0'/0'/0/13 1AN5qBwgrMFYReYq3dGjU29JqWU8Gc9cTE 3AY5WqMcrD9EuZ87XHGG97jKay2uXfs8JK bc1qv667vkm5ew2v7aearlwgz4w4yajlmv4lu54vvx
m/84'/0'/0'/0/14 12LmmGmRD6PWaxwuigNbGWPV2UYpxpZxBq 3Ax1oHbahcAdpgAajuN9khiQA1iAhwSpea bc1qp66k23gwgagadwrsuddzjnsnn0vpewywaw0w8t
m/84'/0'/0'/0/15 1BQ7j6Xwy4RCdx5gPZnruxZie7etpdMJnt 3KSCAVsYYJeWf4GP7KMAWV7yb9ed3JqJ4s bc1qwgg8q70predldmvr0nqma5nxjdkcn2kfzksycn
m/84'/0'/0'/0/16 1GkGsuVefpaaHadaG24MzcZBo7WStrtL5v 32Sw6MyfyWpKJRLHiKrvp96gAkp4yipMnC bc1q4jusduwarl5jsmugtyknx8tv2fpdx52zhd7yw5
m/84'/0'/0'/0/17 1RTu9R7LUFzg7UB16amH6ZbcTvwvZHBbp  38PVZX8THuFvBjqnTh1f6D7nUuZ2C47Bxr bc1qqjsxlud45376fw0c5khhklnkpd3htk0gcms8ad
m/84'/0'/0'/0/18 1M3B2E24exDC1nNXJ6ThpunZ2sHFaHPuwe 36RVnwfPTDmiFMefMCYK6LVMd37nVnidYC bc1qm09mv73f5slu7ms02yrztxmfancwyumqey8rl6
m/84'/0'/0'/0/19 1L5JARNZyXcaxZt1XB2dj6a9K2ug3ZiKUY 37n1Xfx6Jh74MRgvM5EhQ9HCft8XAy3i6e bc1q6yaxz4gr9hrfm8cztmw9eawd7xc8w6mts70p9w


4. xpub
-------

xpriv  xprv9s21ZrQH143K2MPKHPWh91wRxLKehoCNsRrwizj2xNaj9zD5SHMNiHJesDEYgJAavgNE1fDWLgYNneHeSA8oVeVXVYomhP1wxdzZtKsLJbc
xpub   xpub661MyMwAqRbcEqTnPR3hW9tAWNA97FvEEenYXP8eWi7i2nYDypfdG5d8iWfK8YgesKi2EE5mk9THcTqnveDWwZVMuctjmxeEaUKgtg7CEEc
```

## Inspiration

Based on [modood/btckeygen](https://github.com/modood/btckeygen)
