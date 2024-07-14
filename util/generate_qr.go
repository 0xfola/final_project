package util

import (
	"fmt"

	"github.com/skip2/go-qrcode"
)

func GenerateQR(ipfsHash string) error {
	err := qrcode.WriteFile(ipfsHash, qrcode.Medium, 256, "student_id_qr.png")
	if err != nil {
		return fmt.Errorf("error generating QR code: %w", err)
	}
	return nil
}
