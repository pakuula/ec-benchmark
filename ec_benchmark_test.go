package ec_benchmark_test

import (
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"math/big"
	"testing"

	ecrypto "github.com/ethereum/go-ethereum/crypto"
	tmed "github.com/tendermint/tendermint/crypto/ed25519"
	tmsecp "github.com/tendermint/tendermint/crypto/secp256k1"
	tmsr "github.com/tendermint/tendermint/crypto/sr25519"

	"github.com/btcsuite/btcd/btcec/v2"
	btecdsa "github.com/btcsuite/btcd/btcec/v2/ecdsa"
)

var (
	msg  = []byte("Hello, world!")
	hash = sha256.Sum256(msg)
)

func Benchmark_Crypto_Ed25519(b *testing.B) {
	pubkey, privkey, err := ed25519.GenerateKey(rand.Reader)
	dsig := ed25519.Sign(privkey, hash[:])
	if err != nil {
		panic(err)
	}

	for i := 0; i < b.N; i++ {
		hash = sha256.Sum256(msg)
		ed25519.Verify(pubkey, hash[:], dsig)
	}
}

func compress(r, s *big.Int, bitSize int) []byte {
	size := bitSize / 8
	signature := make([]byte, 2*size)
	r.FillBytes(signature[:size])
	s.FillBytes(signature[size:])
	return signature
}

func decompress(signature []byte) (*big.Int, *big.Int) {
	size := len(signature) / 2

	r := big.NewInt(0)
	r.SetBytes(signature[:size])

	s := big.NewInt(0)
	s.SetBytes(signature[size:])

	return r, s
}

func Benchmark_Crypto_P224(b *testing.B) {
	curve := elliptic.P224()
	privkey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		panic(err)
	}
	r, s, err := ecdsa.Sign(rand.Reader, privkey, hash[:])
	if err != nil {
		panic(err)
	}
	signature := compress(r, s, 224)

	for i := 0; i < b.N; i++ {
		r, s := decompress(signature)
		hash = sha256.Sum256(msg)
		if !ecdsa.Verify(&privkey.PublicKey, hash[:], r, s) {
			panic("WTF!")
		}
	}
}

func Benchmark_Crypto_P256(b *testing.B) {
	curve := elliptic.P256()
	privkey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		panic(err)
	}
	r, s, err := ecdsa.Sign(rand.Reader, privkey, hash[:])
	if err != nil {
		panic(err)
	}
	signature := compress(r, s, 256)

	for i := 0; i < b.N; i++ {
		r, s := decompress(signature)
		hash = sha256.Sum256(msg)
		ecdsa.Verify(&privkey.PublicKey, hash[:], r, s)
	}
}

func Benchmark_Crypto_P384(b *testing.B) {
	curve := elliptic.P384()
	privkey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		panic(err)
	}
	r, s, err := ecdsa.Sign(rand.Reader, privkey, hash[:])
	if err != nil {
		panic(err)
	}
	signature := compress(r, s, 384)

	for i := 0; i < b.N; i++ {
		r, s := decompress(signature)
		hash = sha256.Sum256(msg)
		ecdsa.Verify(&privkey.PublicKey, hash[:], r, s)
	}
}

func Benchmark_Crypto_P521(b *testing.B) {
	curve := elliptic.P521()
	privkey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		panic(err)
	}
	r, s, err := ecdsa.Sign(rand.Reader, privkey, hash[:])
	if err != nil {
		panic(err)
	}
	signature := compress(r, s, 528)

	for i := 0; i < b.N; i++ {
		r, s := decompress(signature)
		hash = sha256.Sum256(msg)
		ecdsa.Verify(&privkey.PublicKey, hash[:], r, s)
	}
}

func Must[T any](val T, err error) T {
	if err != nil {
		panic(err)
	}
	return val
}
func Benchmark_Goethereum_Sekp256k1(b *testing.B) {
	privkey := Must(ecrypto.GenerateKey())

	dsig := Must(ecrypto.Sign(hash[:], privkey))
	pubKey := ecrypto.FromECDSAPub(&privkey.PublicKey)
	some := make([]bool, b.N)

	for i := 0; i < b.N; i++ {
		hash = sha256.Sum256(msg)
		some[i] = ecrypto.VerifySignature(pubKey, hash[:], dsig[:64])

	}
}

func Benchmark_Tendermint_Sekp256k1(b *testing.B) {
	privkey := tmsecp.GenPrivKey()
	dsig := Must(privkey.Sign(msg))
	pubKey := privkey.PubKey()
	some := make([]bool, b.N)

	for i := 0; i < b.N; i++ {
		some[i] = pubKey.VerifySignature(msg, dsig)
	}
}

func Benchmark_Tendermint_Ed25519(b *testing.B) {
	privkey := tmed.GenPrivKey()
	dsig := Must(privkey.Sign(msg))
	pubKey := privkey.PubKey()
	some := make([]bool, b.N)

	for i := 0; i < b.N; i++ {
		some[i] = pubKey.VerifySignature(msg, dsig)
		if !some[i] {
			panic("WTF!")
		}
	}
}

func Benchmark_Tendermint_Sr25519(b *testing.B) {
	privkey := tmsr.GenPrivKey()
	dsig := Must(privkey.Sign(msg))
	pubKey := privkey.PubKey()
	some := make([]bool, b.N)

	for i := 0; i < b.N; i++ {
		some[i] = pubKey.VerifySignature(msg, dsig)
	}
}

func Benchmark_Btcsuite_ECDSA(b *testing.B) {
	// Decode a hex-encoded private key.
	pkBytes := Must(hex.DecodeString("22a47fa09a223f2aa079edf85a7c2d4f87" +
		"20ee63e502ee2869afab7de234b80c"))
	privKey, pubKey := btcec.PrivKeyFromBytes(pkBytes)
	dsig := btecdsa.Sign(privKey, hash[:])
	some := make([]bool, b.N)

	for i := 0; i < b.N; i++ {
		some[i] = dsig.Verify(hash[:], pubKey)
	}
}
