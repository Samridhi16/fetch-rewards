package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fetch-rewards/models"
	"strconv"
)

func GenerateUniqueID(r models.Receipt) string {
	//concatenates receipt's details in order to make sure unique id is created for unique receipt
	concatenated := r.Retailer + "-" + r.PurchaseDate + "-" + r.PurchaseTime + "-" + r.Total + "-" + strconv.Itoa(len(r.Items))
	// Generate SHA256 hash of the concatenated string
	hash := sha256.Sum256([]byte(concatenated))
	// Return the hexadecimal representation of the hash
	return hex.EncodeToString(hash[:])
}
