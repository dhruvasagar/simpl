use anyhow::Result;
use sqlx::{Pool, Sqlite};

pub struct Report {}

impl Report {
    pub async fn merchant_discount(merchant_name: String, pool: &Pool<Sqlite>) -> Result<f64> {
        let discount: (f64,) = sqlx::query_as(
            "SELECT sum(amount) * merchants.discount_percentage / 100 as discount
            FROM transactions
            INNER JOIN merchants on merchants.id = transactions.merchant_id
            WHERE merchant_id = (SELECT id FROM merchants WHERE name = $1)
            ",
        )
            .bind(merchant_name)
            .fetch_one(pool)
            .await?;
        Ok(discount.0)
    }

    pub async fn user_dues(user_name: String, pool: &Pool<Sqlite>) -> Result<f64> {
		let dues: (f64,) = sqlx::query_as(
            "SELECT sum(amount) as total_amount
            FROM transactions
            WHERE user_id = (SELECT id FROM users WHERE name = $1)
            ",
        )
            .bind(user_name)
            .fetch_one(pool)
            .await?;
        Ok(dues.0)
    }

    pub async fn users_at_credit_limit(pool: &Pool<Sqlite>) -> Result<Vec<String>> {
        let users: Vec<(String,)> = sqlx::query_as(
            "SELECT users.name
            FROM transactions
            INNER JOIN users on users.id = transactions.user_id
            GROUP BY user_id
            HAVING sum(amount) = users.credit_limit
            ",
        )
            .fetch_all(pool)
            .await?;
        Ok(users.iter().map(|u| u.clone().0).collect())
    }

    pub async fn users_total_dues(pool: &Pool<Sqlite>) -> Result<Vec<(String, f64)>> {
        let users: Vec<(String, f64)> = sqlx::query_as(
            "SELECT users.name, sum(transactions.amount) as due
            FROM users
            LEFT JOIN transactions on users.id = transactions.user_id
            GROUP BY user_id
            UNION ALL
            SELECT 'total', sum(amount) as due
            FROM transactions
            ",
        )
            .fetch_all(pool)
            .await?;
        Ok(users)
    }
}
