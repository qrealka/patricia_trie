# Intro

This is a simplified implementation of Ethereum's modified Merkle Patricia Trie 

## API

Ethereum's Merkle Patricia Trie is essentially a key-value mapping that provides the following standard methods:

```go
type Trie interface {
  // methods as a basic key-value mapping
  Get(key []byte) ([]byte, bool) {
  Put(key []byte, value []byte)
  Del(key []byte, value []byte) bool
}
```

## Original code

Full source code can be found [here](https://github.com/zhangchiqing/merkle-patricia-trie/tree/master)
