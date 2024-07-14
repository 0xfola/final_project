package server

import (
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

func HttpServer() {
	r := gin.Default()

	r.POST("/store_student_data", func(c *gin.Context) {
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

	r.GET("/get_student_data/:studentAddress", func(c *gin.Context) {
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
