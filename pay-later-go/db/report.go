package db

import "database/sql"

// MerchantDiscount computes the total discount for a merchant
func (db *DB) MerchantDiscount(merchant Merchant) (float64, error) {
	var discount sql.NullFloat64
	db.db.Raw(
		`
		SELECT sum(amount) * merchants.discount_percentage / 100 as discount
		FROM transactions
		INNER JOIN merchants on merchants.id = transactions.merchant_id
		WHERE merchant_id = ?
		`,
		merchant.ID,
	).Scan(&discount)
	return discount.Float64, nil
}

// UserDues computes the total dues of a user
func (db *DB) UserDues(user User) float64 {
	var userTotalDues sql.NullFloat64
	db.db.Raw(
		`
		SELECT sum(amount) as total_amount
		FROM transactions
		WHERE user_id = ?
		`,
		user.ID,
	).Scan(&userTotalDues)
	return userTotalDues.Float64
}

// UsersAtCreditLimit lists all users who have reached their credit limit
func (db *DB) UsersAtCreditLimit() []User {
	var users []User
	db.db.Raw(
		`
		SELECT users.*
		FROM transactions
		INNER JOIN users on users.id = transactions.user_id
		GROUP BY user_id
		HAVING sum(amount) = users.credit_limit
		`,
	).Scan(&users)
	return users
}

// UserDue encapsulates each users name and their total due amount
type UserDue struct {
	Name string
	Due  float64
}

// UsersTotalDues lists all users and their total due amount
// Additionally this also includes the total sum of all dues by all users
func (db *DB) UsersTotalDues() []UserDue {
	var userDues []UserDue
	db.db.Raw(
		`
		SELECT users.name, sum(amount) as due
		FROM transactions
		INNER JOIN users on users.id = transactions.user_id
		GROUP BY user_id
		UNION ALL
		SELECT 'total', sum(amount) as due
		FROM transactions
		`,
	).Scan(&userDues)
	return userDues
}
