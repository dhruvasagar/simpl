use anyhow::{anyhow, Result};
use sqlx::{Pool, Sqlite, FromRow};
use crate::db::{User, Merchant, Report};

#[derive(Debug, FromRow)]
pub struct Transaction {
    pub id: Option<i64>,
    pub user_id: i64,
    pub merchant_id: Option<i64>,
    pub amount: f64,
}

impl Transaction {
    #[allow(dead_code)]
    pub fn new(user_id: i64, merchant_id: Option<i64>, amount: f64) -> Self {
        Self { id: None, user_id, merchant_id, amount }
    }

    pub async fn create(user: User, merchant: Merchant, amount: f64, pool: &Pool<Sqlite>) -> Result<()> {
        let user_dues = Report::user_dues(user.name, pool).await?;
        if user_dues+amount > user.credit_limit {
            return Err(anyhow!("credit limit"));
        }

        sqlx::query("INSERT INTO transactions (user_id, merchant_id, amount) VALUES ($1, $2, $3)")
            .bind(user.id)
            .bind(merchant.id)
            .bind(amount)
            .execute(pool)
            .await?;
        Ok(())
    }

    pub async fn payback(user: User, amount: f64, pool: &Pool<Sqlite>) -> Result<()> {
        sqlx::query("INSERT INTO transactions (user_id, amount) VALUES ($1, $2)")
            .bind(user.id)
            .bind(-1.0 * amount)
            .execute(pool)
            .await?;
        Ok(())
    }
}
