package main

import (
	"final-project/util"
	"fmt"
	"strings"

	shell "github.com/ipfs/go-ipfs-api"
)

func main() {
	studentData := []byte(`{"name": "John Doe", "age": 20, "student_id": "12345", "image": "base64_encoded_image_data"}`)
	passphrase := "Sixteen byte key"

	// Encrypt student data
	ciphertext, nonce, tag := util.Encrypt(studentData, passphrase)
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
