package main

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "encoding/hex"
    "fmt"
    "io"
    "strings"

    "github.com/ipfs/go-ipfs-api"
)

// encrypt encrypts data using AES-GCM with the given passphrase
func encrypt(data []byte, passphrase string) (string, string, string) {
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

func main() {
    studentData := []byte(`{"name": "John Doe", "age": 20, "student_id": "12345", "image": "base64_encoded_image_data"}`)
    passphrase := "Sixteen byte key"

    // Encrypt student data
    ciphertext, nonce, tag := encrypt(studentData, passphrase)
    fmt.Println("Ciphertext:", ciphertext)
    fmt.Println("Nonce:", nonce)
    fmt.Println("Tag:", tag)

    // Connect to IPFS
    sh := shell.NewShell("localhost:5001")

    // Upload encrypted data to IPFS
    cid, err := sh.Add(strings.NewReader(ciphertext))
    if err != nil {
        fmt.Println("Error uploading to IPFS:", err)
        return
    }
    fmt.Println("IPFS CID:", cid)
}