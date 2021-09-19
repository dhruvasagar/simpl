package db

import (
	"fmt"

	"github.com/dhruvasagar/pay-later-go/utils"
	"gorm.io/gorm"
)

// Merchant is a gorm database model representing the merchants data
type Merchant struct {
	gorm.Model
	Name               string
	DiscountPercentage float64
}

func (m Merchant) String() string {
	return fmt.Sprintf("%s(%.4s%%)", m.Name, utils.FormatFloat(m.DiscountPercentage))
}

// CreateMerchant inserts a merchant record in the database
func (db *DB) CreateMerchant(merchant Merchant) (Merchant, error) {
	result := db.db.Create(&merchant)
	return merchant, result.Error
}

// UpdateMerchant updates a merchant record in the database
// NOTE: ID must be present on the model, otherwise this inserts a new record
//       instead of updating
func (db *DB) UpdateMerchant(merchant Merchant) (Merchant, error) {
	result := db.db.Save(&merchant)
	return merchant, result.Error
}

// FindMerchantByName looks up a merchant record by name from the database
func (db *DB) FindMerchantByName(name string) (Merchant, error) {
	var merchant Merchant
	result := db.db.Where("name = ?", name).First(&merchant)
	return merchant, result.Error
}
