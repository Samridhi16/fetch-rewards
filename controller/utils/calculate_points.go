package utils

import (
	"fetch-rewards/models"
	"math"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func CalculatePoints(r models.Receipt) int64 {
	var points int64 = 0

	// Rewards one point for every alphanumeric character in the retailer name.
	for _, char := range r.Retailer {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			points++
		}
	}

	totalFloat, err := strconv.ParseFloat(r.Total, 64)
	if err == nil {
		//Rewards 50 points if the total is a round dollar amount with no cents.
		if math.Mod(totalFloat, 1) == 0 {
			points += 50
		}

		// Rewards 25 points if the total is a multiple of 0.25.
		if math.Mod(totalFloat, 0.25) == 0 {
			points += 25
		}
	}

	// Rewards 5 points for every two items on the receipt.
	points += (int64)(len(r.Items)/2) * 5

	//If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and
	//round up to the nearest integer. The result is the number of points earned.
	for _, i := range r.Items {
		trimmedDescription := strings.TrimSpace(i.ShortDescription)
		if len(trimmedDescription)%3 == 0 {
			itemPrice, _ := strconv.ParseFloat(i.Price, 64)
			points += int64(math.Ceil(itemPrice * 0.2))
		}
	}

	// Rewards 6 points if the day in the purchase date is odd.
	purchaseDate, _ := time.Parse("2006-01-02", r.PurchaseDate)
	if purchaseDate.Day()%2 != 0 {
		points += 6
	}

	// Rewards 10 points if the time of purchase is after 2:00pm and before 4:00pm.
	purchaseTime, _ := time.Parse("15:04", r.PurchaseTime)
	if purchaseTime.Hour() >= 14 && purchaseTime.Hour() < 16 {
		points += 10
	}

	return points
}
