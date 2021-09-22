use std::fmt;
use anyhow::Result;
use sqlx::{FromRow, Pool, Sqlite};

#[derive(Debug, FromRow, Clone)]
pub struct Merchant {
    pub id: Option<i64>,
    pub name: String,
    pub discount_percentage: f64,
}

impl fmt::Display for Merchant {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "{}({})", self.name, self.discount_percentage)
    }
}

impl Merchant {
    pub fn new(name: String, discount_percentage: f64) -> Self {
        Self {
            id: None,
            name,
            discount_percentage,
        }
    }

    pub async fn create(merchant: Merchant, pool: &Pool<Sqlite>) -> Result<Merchant> {
        let mut rmerchant = merchant.clone();
        let merchant_id: (i64,) = sqlx::query_as("INSERT INTO merchants (name, discount_percentage) VALUES ($1, $2) RETURNING id")
            .bind(merchant.name)
            .bind(merchant.discount_percentage)
            .fetch_one(pool)
            .await?;
        rmerchant.id = Some(merchant_id.0);
        Ok(rmerchant)
    }

    pub async fn update(merchant: Merchant, pool: &Pool<Sqlite>) -> Result<()> {
        sqlx::query("UPDATE merchants SET discount_percentage=$1 WHERE name=$2")
            .bind(merchant.discount_percentage)
            .bind(merchant.name)
            .execute(pool)
            .await?;
        Ok(())
    }

    pub async fn find_by_name(merchant_name: String, pool: &Pool<Sqlite>) -> Result<Merchant> {
        let merchant = sqlx::query_as("SELECT * FROM merchants WHERE name=$1")
            .bind(merchant_name)
            .fetch_one(pool)
            .await?;
        Ok(merchant)
    }
}
