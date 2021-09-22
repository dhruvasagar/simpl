-- Add migration script here
CREATE TABLE IF NOT EXISTS transactions (
  id INTEGER NOT NULL PRIMARY KEY,
  user_id INTEGER NOT NULL,
  merchant_id INTEGER,
  amount REAL,

  FOREIGN KEY(user_id) REFERENCES users(id),
  FOREIGN KEY(merchant_id) REFERENCES merchants(id)
);
