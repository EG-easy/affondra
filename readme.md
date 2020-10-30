# affondra

<p align="center">
  <img src="./affondra-logo.jpg" width="300">
</p>

[![version](https://img.shields.io/github/v/tag/EG-easy/affondra)](https://github.com/EG-easy/affondra/releases/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/EG-easy/affondra)](https://goreportcard.com/report/github.com/EG-easy/affondra)

Affondra is a blockchain application using cosmos-sdk that anyone can use when trading NFT items.

# Implementation
## Completed
- [x] basic blockchain functions
- [x] issue NFT token
- [x] create Item based on NFT
- [x] update Item infomation
- [x] buy Item
- [x] affiliate system
- [x] faucet system
- [x] web wallet

## Current Work

## Next Steps

# Quick Start
Go 1.15+ is required for the Cosmos SDK.

Firstly, watch [Demo]()!

## Install affondracli and affondrad

```bash
$ git clone github.com/EG-easy/affondra.git
$ cd affondra && git checkout main
$ make
```

Try `votumcli version` and `affondrad version` to verify everything is OK!

## **Initialize configuration files and genesis file**

Just use shell scripts bellow.
```
$ sh scripts/start.sh
```

**Now let's start!**
```bash
$ affondrad start
```

## Functions

### Mint NFT
Firstly, you need to mint(or receive) a NFT.
When creating NFT, you can modify the denom, nft id, description, tokenURI.

```bash
$ affondracli tx nft mint your-favorite-denom random-nft-id $(affondracli keys show -a user1) --tokenURI http://metadata.com --from user1 -y
```

### Create Item
Now you have a NFT that you can sell the NFT as a Item in the Affondra trading market. Just set the price you want to sell and affiliate reward for those who promote your item.

```bash
$ affondracli tx affondra create-item your-favorite-denom random-nft-id 1000affondollar 10affondollar This item is Awesome! true --from=user1 -y
```

### Buy Item
Now your friend A introduce the item to your friend B, and friend B bought the item with the command below. The result will be you get
990affondollar and your friend gets 10affondollar as a affiliate reward, and B will successfully receive NFT from you.
```bash
$ votumcli tx affondra buy-item random-nft-id B'address A'address --from B -y
```

## License
Licensed under the [MIT](LICENSE).

