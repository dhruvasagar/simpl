package db

import (
	"errors"

	"gorm.io/gorm"
)

// Transaction is a gorm database model represent the transactions data
type Transaction struct {
	gorm.Model
	UserID     uint
	MerchantID uint
	Amount     float64

	User     User
	Merchant Merchant
}

// CreditLimit is the error message when rejecting a transaction when
// processing the current transaction will result in his total credit amount
// to exceed his credit limit
const CreditLimit = "credit limit"

// CreateTransaction inserts a transaction record in the database
func (db *DB) CreateTransaction(
	user User,
	merchant Merchant,
	amount float64,
) (Transaction, error) {
	userTotalCredit := db.UserDues(user)
	if userTotalCredit+amount > user.CreditLimit {
		return Transaction{}, errors.New(CreditLimit)
	}

	txn := Transaction{
		UserID:     user.ID,
		MerchantID: merchant.ID,
		Amount:     amount,
	}
	result := db.db.Create(&txn)
	return txn, result.Error
}

// CreatePayback inserts a transaction record in the database as payback for
// a user. This does not correspond to a particular merchant since it's received
// by our service and so the merchant_id for this transaction will be 0.
func (db *DB) CreatePayback(
	user User,
	amount float64,
) (Transaction, error) {
	txn := Transaction{
		UserID: user.ID,
		Amount: -1 * amount,
	}
	result := db.db.Create(&txn)
	return txn, result.Error
}
