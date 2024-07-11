package util

import (
    "fmt"
    "github.com/skip2/go-qrcode"
)

func GenerateQR() {
    ipfsHash := "QmbTq7zHWX3hDez4iqEqJBna3RQwHZuymNuY8jFRVjUZ1J" // Replace with your IPFS hash

    // Generate QR code
    err := qrcode.WriteFile(ipfsHash, qrcode.Medium, 256, "student_id_qr.png")
    if err != nil {
        fmt.Println("Error generating QR code:", err)
    }
}
