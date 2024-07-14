package main

import (
	"final-project/util"
	"strings"

	shell "github.com/ipfs/go-ipfs-api"

	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/blocto/solana-go-sdk/client"
	"github.com/blocto/solana-go-sdk/rpc"
	"github.com/gin-gonic/gin"
)

type StoreStudentDataRequest struct {
	StudentAddress string `json:"studentAddress"`
	IpfsHash       string `json:"ipfsHash"`
	Nonce          string `json:"nonce"`
	Tag            string `json:"tag"`
}

func main() {
	r := gin.Default()

	r.POST("/student", func(c *gin.Context) {
		var req StoreStudentDataRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Connect to Solana client
		cli := client.NewClient(rpc.DevnetRPCEndpoint)

		txHash, err := cli.GetLatestBlockhash(context.Background())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		fmt.Println("Transaction Hash:", txHash)

		c.JSON(http.StatusOK, gin.H{"message": "Student data stored successfully.", "txHash": txHash})
	})

	r.GET("/student/:studentAddress", func(c *gin.Context) {
		studentAddress := c.Param("studentAddress")

		// Connect to Solana client
		cli := client.NewClient(rpc.DevnetRPCEndpoint)

		// Note: Implement the actual logic to retrieve student data from Solana
		accountInfo, err := cli.GetAccountInfo(context.Background(), studentAddress)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		fmt.Println("Account Info:", accountInfo)

		c.JSON(http.StatusOK, gin.H{"studentAddress": studentAddress, "accountInfo": accountInfo})
	})

	if err := r.Run(":3000"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

func test() {
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
