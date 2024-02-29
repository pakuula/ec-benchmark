Benchmarks of several ECDSA implementions.

# Run
```bash
make run
```

# Benchmarks

**Benchmark_Crypto_Ed25519** benchmark verify signature for Ed25519 curve from the standard library
[`crypto/ed25519`](https://pkg.go.dev/crypto/ed25519)

**Benchmark_Crypto_P224** benchmark verify signature for NIST P-224 curve from the standard library
[`crypto/elliptic`](https://pkg.go.dev/crypto/elliptic)

**Benchmark_Crypto_P256** benchmark verify signature for NIST P-256 curve from the standard library
[`crypto/elliptic`](https://pkg.go.dev/crypto/elliptic)

**Benchmark_Crypto_P384** benchmark verify signature for NIST P-384 curve from the standard library
[`crypto/elliptic`](https://pkg.go.dev/crypto/elliptic)

**Benchmark_Crypto_P521** benchmark verify signature for NIST P-521 curve from the standard library
[`crypto/elliptic`](https://pkg.go.dev/crypto/elliptic)

**Benchmark_Goethereum_Sekp256k1** benchmark verify signature for secp256k1 curve from 
[`go-ethereum`](https://pkg.go.dev/github.com/ethereum/go-ethereum/crypto)

**Benchmark_Tendermint_Sekp256k1** benchmark verify signature for secp256k1 curve from 
[Tendermint `crypto/secp256k1`](https://pkg.go.dev/github.com/tendermint/tendermint/crypto/secp256k1)

**Benchmark_Tendermint_Ed25519** benchmark verify signature for Ed25519 curve from 
[Tendermint `crypto/ed25519`](https://pkg.go.dev/github.com/tendermint/tendermint/crypto/ed25519)

**Benchmark_Tendermint_Sr25519** benchmark verify signature for Sr25519 curve from 
[Tendermint `crypto/sr25519`](https://pkg.go.dev/github.com/tendermint/tendermint/crypto/sr25519) 
(Schnorrkel/Ristretto x25519 signature scheme)

**Benchmark_Btcsuite_ECDSA** benchmark verify signature for secp256k1 curve from 
[BTCD ECDSA](https://pkg.go.dev/github.com/btcsuite/btcd/btcec/v2/ecdsa)

# Sample results

Results as of 2024-02-29:

```text
go test -benchmem -bench .
goos: linux
goarch: amd64
pkg: github.com/pakuula/ec_benchmark
cpu: 12th Gen Intel(R) Core(TM) i9-12900
Benchmark_Crypto_Ed25519-24                40016             29882 ns/op               0 B/op          0 allocs/op
Benchmark_Crypto_P224-24                    9920            119372 ns/op            1104 B/op         28 allocs/op
Benchmark_Crypto_P256-24                   23425             51222 ns/op            1136 B/op         24 allocs/op
Benchmark_Crypto_P384-24                    2966            390831 ns/op            1409 B/op         28 allocs/op
Benchmark_Crypto_P521-24                    1167           1013839 ns/op            1908 B/op         29 allocs/op
Benchmark_Goethereum_Sekp256k1-24          40644             29554 ns/op               1 B/op          0 allocs/op
Benchmark_Tendermint_Sekp256k1-24           6968            168197 ns/op            4184 B/op         88 allocs/op
Benchmark_Tendermint_Ed25519-24            68068             18139 ns/op              81 B/op          2 allocs/op
Benchmark_Tendermint_Sr25519-24            48381             24189 ns/op             473 B/op          5 allocs/op
Benchmark_Btcsuite_ECDSA-24                10000            107722 ns/op            1410 B/op         27 allocs/op
```