-- Add migration script here
CREATE TABLE IF NOT EXISTS merchants (
  id INTEGER NOT NULL PRIMARY KEY,
  name TEXT NOT NULL UNIQUE,
  discount_percentage REAL
);
