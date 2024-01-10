package modelutils

import (
	"fetch-rewards/models"
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"
)

var (
	// retailerRegex         = regexp.MustCompile("^\\S+$")
	totalRegex            = regexp.MustCompile("^\\d+\\.\\d{2}$")
	shortDescriptionRegex = regexp.MustCompile("^[\\w\\s\\-]+$")
	priceRegex            = regexp.MustCompile("^\\d+\\.\\d{2}$")
)

func ValidateReceipt(r models.Receipt) error {
	validate := validator.New()
	if err := validate.Struct(r); err != nil {
		return err
	}
	return validateRegex(r)
}

func validateRegex(r models.Receipt) error {

	//Provided Regex invalid for "M&M Corner Market" hence commented out
	// if !retailerRegex.MatchString(r.Retailer) {
	// 	return fmt.Errorf("Invalid Retailer name")
	// }

	if !totalRegex.MatchString(r.Total) {
		return fmt.Errorf("Invalid Total")
	}

	for i, item := range r.Items {
		if !shortDescriptionRegex.MatchString(item.ShortDescription) {
			return fmt.Errorf("Invalid ShortDecription for item %d", i)
		}
		if !priceRegex.MatchString(item.Price) {
			return fmt.Errorf("Invalid Price for item %d", i)
		}
	}
	return nil
}
