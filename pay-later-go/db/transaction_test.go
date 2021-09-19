package db

import (
	"os"
	"testing"
)

func TestCreateTransaction(t *testing.T) {
	os.Setenv("BOLT_PATH", "file::memory:?cache=shared")
	sdb := New()

	user := User{
		Name:        "test_user",
		Email:       "test_email@test.com",
		CreditLimit: 100,
	}
	user, _ = sdb.CreateUser(user)

	merchant := Merchant{
		Name:               "test_merchant",
		DiscountPercentage: 1.25,
	}
	merchant, _ = sdb.CreateMerchant(merchant)

	_, err := sdb.CreateTransaction(user, merchant, 100)
	if err != nil {
		t.Fatal("Failed to create transaction")
	}

	_, err = sdb.CreateTransaction(user, merchant, 100)
	if err.Error() != "credit limit" {
		t.Fatal("Failed to reject properly because of credit limit")
	}
	if err == nil {
		t.Fatal("Failed to reject transaction")
	}
}
