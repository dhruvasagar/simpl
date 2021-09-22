use std::env::var;
use anyhow::Result;
use sqlx::{migrate, sqlite::SqlitePoolOptions, Pool, Sqlite};
use crate::db::{User, Merchant, Transaction, Report};

pub struct DB {
    pub pool: Pool<Sqlite>,
}

impl DB {
    pub async fn new() -> Result<Self> {
        let db_path = match var("DB_PATH") {
            Ok(p) => p,
            Err(_) => "pay-later-rs.db".to_string(),
        };

        let pool = SqlitePoolOptions::new()
            .connect(db_path.as_str())
            .await?;

        let migrator = migrate!();
        migrator.run(&pool).await?;

        Ok(Self { pool })
    }

    pub async fn create_user(&self, user: User) -> Result<User> {
        User::create(user, &self.pool.clone()).await
    }

    pub async fn update_user(&self, user: User) -> Result<()> {
        User::update(user, &self.pool.clone()).await
    }

    pub async fn find_user(&self, user_name: String) -> Result<User> {
        User::find_by_name(user_name, &self.pool.clone()).await
    }

    pub async fn create_merchant(&self, merchant: Merchant) -> Result<Merchant> {
        Merchant::create(merchant, &self.pool.clone()).await
    }

    pub async fn update_merchant(&self, merchant: Merchant) -> Result<()> {
        Merchant::update(merchant, &self.pool.clone()).await
    }

    pub async fn find_merchant(&self, merchant_name: String) -> Result<Merchant> {
        Merchant::find_by_name(merchant_name, &self.pool.clone()).await
    }

    pub async fn create_transaction(&self, user: User, merchant: Merchant, amount: f64) -> Result<()> {
        Transaction::create(user, merchant, amount, &self.pool.clone()).await
    }

    pub async fn create_payback(&self, user: User, amount: f64) -> Result<()> {
        Transaction::payback(user, amount, &self.pool.clone()).await
    }

    pub async fn report_discount(&self, merchant_name: String) -> Result<f64> {
        Report::merchant_discount(merchant_name, &self.pool.clone()).await
    }

    pub async fn report_user_dues(&self, user_name: String) -> Result<f64> {
        Report::user_dues(user_name, &self.pool.clone()).await
    }

    pub async fn report_users_at_credit_limit(&self) -> Result<Vec<String>> {
        Report::users_at_credit_limit(&self.pool.clone()).await
    }

    pub async fn report_users_total_dues(&self) -> Result<Vec<(String, f64)>> {
        Report::users_total_dues(&self.pool.clone()).await
    }
}
