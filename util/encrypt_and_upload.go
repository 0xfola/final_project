package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"io"
)

// encrypt encrypts data using AES-GCM with the given passphrase
func Encrypt(data []byte, passphrase string) (string, string, string) {
	key := []byte(passphrase)
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return hex.EncodeToString(ciphertext), hex.EncodeToString(nonce), hex.EncodeToString(gcm.Seal(nil, nonce, nil, nil))
}
