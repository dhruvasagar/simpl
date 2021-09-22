use std::fmt;
use anyhow::Result;
use sqlx::{Pool, Sqlite, FromRow};

#[derive(Debug, FromRow, Clone)]
pub struct User {
    pub id: Option<i64>,
    pub name: String,
    pub email: String,
    pub credit_limit: f64,
}

impl fmt::Display for User {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "{}({})", self.name, self.credit_limit)
    }
}

impl User {
    pub fn new(name: String, email: String, credit_limit: f64) -> Self {
        Self {
            id: None,
            name,
            email,
            credit_limit,
        }
    }

    pub async fn create(user: User, pool: &Pool<Sqlite>) -> Result<User> {
        let mut ruser = user.clone();
        let user_id: (i64,) = sqlx::query_as("INSERT INTO users (name, email, credit_limit) VALUES ($1, $2, $3) RETURNING id")
            .bind(user.name)
            .bind(user.email)
            .bind(user.credit_limit)
            .fetch_one(pool)
            .await?;
        ruser.id = Some(user_id.0);
        Ok(ruser)
    }

    pub async fn update(user: User, pool: &Pool<Sqlite>) -> Result<()> {
        sqlx::query("UPDATE users SET email=$1, credit_limit=$2 WHERE name=$3")
            .bind(user.email)
            .bind(user.credit_limit)
            .bind(user.name)
            .execute(pool)
            .await?;
        Ok(())
    }

    pub async fn find_by_name(user_name: String, pool: &Pool<Sqlite>) -> Result<User> {
        let user = sqlx::query_as("SELECT * FROM users WHERE name=$1")
            .bind(user_name)
            .fetch_one(pool)
            .await?;
        Ok(user)
    }
}
