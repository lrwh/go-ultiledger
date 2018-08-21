package crypto

import (
	"crypto/rand"
	"errors"
	"fmt"
	"io"

	b58 "github.com/mr-tron/base58/base58"
	"golang.org/x/crypto/ed25519"
)

// generate key pair with ed25510 cryto algorithm, since we can
// always reconstruct the true private key using the same seed,
// we use the randomly generated seed as a equivalent private key.
func ed25519Keypair() (string, string, error) {
	var seed [32]byte
	_, err := io.ReadFull(rand.Reader, seed[:])
	if err != nil {
		return "", "", err
	}
	privateKey := ed25519.NewKeyFromSeed(seed[:])
	publicKey := privateKey.Public().(ed25519.PublicKey)

	var pk [32]byte
	copy(pk[:], publicKey)
	acc := &ULTKey{Code: KeyTypeAccountID, Hash: pk}
	sd := &ULTKey{Code: KeyTypeSeed, Hash: seed}

	pubKeyStr := EncodeKey(acc)
	seedStr := EncodeKey(sd)

	return pubKeyStr, seedStr, nil
}

// reconstruct the true private key from the seed, it supposes to
// be only used in situations where you need to sign the data so
// the authenticity can be verified by the corresponding public key.
func getPrivateKey(seed string) (ed25519.PrivateKey, error) {
	if seed == "" {
		return nil, fmt.Errorf("empty seed")
	}
	k, err := DecodeKey(seed)
	if err != nil {
		return nil, err
	}
	privateKey := ed25519.NewKeyFromSeed(k.Hash[:])
	return privateKey, nil
}

// randomly generate a pair of public and private key
func GenerateKeypair() (string, string, error) {
	// privateKey is actually the seed used to generate the keypair
	publicKey, seed, err := ed25519Keypair()
	if err != nil {
		return "", "", err
	}
	return publicKey, seed, err
}

func GenerateKeypairFromSeed(seed []byte) (string, string, error) {
	if len(seed) != 32 {
		return "", "", errors.New("Invalid seed, byte length is not 32")
	}
	privateKey := ed25519.NewKeyFromSeed(seed)
	publicKey := privateKey.Public().(ed25519.PublicKey)

	var pk [32]byte
	copy(pk[:], publicKey)
	acc := &ULTKey{Code: KeyTypeAccountID, Hash: pk}

	var sdk [32]byte
	copy(sdk[:], seed)
	sd := &ULTKey{Code: KeyTypeSeed, Hash: sdk}

	pubKeyStr := EncodeKey(acc)
	seedStr := EncodeKey(sd)

	return pubKeyStr, seedStr, nil

}

// sign the data with provided seed (equivalent private key)
func Sign(seed string, data []byte) (string, error) {
	pk, err := getPrivateKey(seed)
	if err != nil {
		return "", err
	}

	signature := ed25519.Sign(pk, data)
	signStr := b58.Encode(signature)

	return signStr, nil
}

// verify the data signature
func Verify(publicKey, signature string, data []byte) bool {
	pk, err := DecodeKey(publicKey)
	if err != nil {
		return false
	}
	sn, err := b58.Decode(signature)
	if err != nil {
		return false
	}
	pub := ed25519.PublicKey(pk.Hash[:])
	return ed25519.Verify(pub, data, sn)
}
