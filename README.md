# This Project follows the following tutorial

- Currently on Blockchain in Go Part 6
- You can find links for whole tutorial [here](https://github.com/practical-tutorials/project-based-learning#go)

## Repo created at midway of Part 4

- Created this to learn go and implementation of block chain, base58 algorithms I yoinked it.

# TODOS

- [x] Send crypto to other address
- [x] Private key addresses
- [x] Wallets
- [x] Transaction Signature
- [ ] Expand each block's max transactions record (currently: 1 / block)
- [x] Improve transaction (e.g. Rewards, Merkle Tree)
- [ ] Networking

# Build this noob app

```
$ go build -o blockchain_go *.go
```

# Run after building

```
$ ./blockchain_go
Usage:
  getbalance -address ADDRESS - Get Balance of ADDRESS
  printchain - print all the blocks of the blockchain
  send -from FROM -to TO -amount AMOUNT - Send AMOUNT of coins from FROM address to TO
```

# Commit Log

- All of them is me, my ssh setup is weird and I am testing it
