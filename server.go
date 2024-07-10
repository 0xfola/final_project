package main

import (
    "context"
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/gagliardetto/solana-go"
    "log"
    "net/http"
)

type StoreStudentDataRequest struct {
    StudentAddress string `json:"studentAddress"`
    IpfsHash       string `json:"ipfsHash"`
    Nonce          string `json:"nonce"`
    Tag            string `json:"tag"`
}

func main() {
    r := gin.Default()

    r.POST("/store_student_data", func(c *gin.Context) {
        var req StoreStudentDataRequest
        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        // Connect to Solana client
        cli := client.NewClient(rpc.DevnetRPCEndpoint)

        // Example of sending a transaction
        // Note: Implement the actual logic to store student data on Solana
        txHash, err := cli.GetRecentBlockhash(context.Background())
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

        // Example of retrieving account data
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