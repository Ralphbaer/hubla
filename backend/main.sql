-- Table to store the types of transactions
CREATE TYPE transaction_type AS ENUM ('CREATOR_SALE', 'AFFILIATE_SALE', 'COMMISSION_PAID', 'COMMISSION_RECEIVED');

CREATE TYPE seller_type AS ENUM ('CREATOR', 'AFFILIATE');

CREATE TABLE file_metadata (
    id UUID PRIMARY KEY,
    file_size BIGINT NOT NULL,
    disposition TEXT NOT NULL,
    hash TEXT NOT NULL UNIQUE,
    binary_data BYTEA NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);

-- Table to store the sellers
CREATE TABLE sellers (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    seller_type seller_type NOT NULL, -- new column to identify the seller role
    created_at TIMESTAMP NOT NULL DEFAULT now()
);

-- Table to store the products
CREATE TABLE products (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    creator_id UUID NOT NULL REFERENCES sellers(id),
    created_at TIMESTAMP NOT NULL DEFAULT now()
);

-- Table to store the transactions
CREATE TABLE transactions (
    id UUID PRIMARY KEY,
    t_type transaction_type NOT NULL,
    t_date TIMESTAMP NOT NULL,
	product_id UUID NOT NULL REFERENCES products(id),
    amount NUMERIC(19, 2) NOT NULL,
    seller_id UUID NOT NULL REFERENCES sellers(id),
    created_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE file_transactions (
	id UUID PRIMARY KEY,
	file_id UUID NOT NULL,
	transaction_id UUID NOT NULL,
	FOREIGN KEY (file_id) REFERENCES file_metadata(id),
	FOREIGN KEY (transaction_id) REFERENCES transactions(id),
	UNIQUE (file_id, transaction_id)
);

-- Table to store the seller balances
CREATE TABLE seller_balances (
    id UUID PRIMARY KEY,
    seller_id UUID NOT NULL REFERENCES sellers(id),
    balance NUMERIC(19, 2) NOT NULL,
	updated_at TIMESTAMP NOT NULL DEFAULT now(),
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    UNIQUE(seller_id)
);

CREATE TABLE users (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE EXTENSION IF NOT EXISTS "pgcrypto";
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE OR REPLACE FUNCTION hash_password(password TEXT)
RETURNS TEXT AS $$
BEGIN
    RETURN encode(pgcrypto.digest(password, 'sha512'), 'hex');
END;
$$ LANGUAGE plpgsql;


CREATE INDEX idx_transactions_seller_id ON transactions(seller_id);
CREATE INDEX idx_seller_balances_seller_id ON seller_balances(seller_id);

TRUNCATE TABLE SELLERS CASCADE;
TRUNCATE TABLE TRANSACTIONS CASCADE;
TRUNCATE TABLE PRODUCTS CASCADE;
TRUNCATE TABLE SELLER_BALANCES CASCADE;
TRUNCATE TABLE FILE_METADATA CASCADE;
TRUNCATE TABLE FILE_TRANSACTIONS CASCADE;

DROP TABLE IF EXISTS SELLERS CASCADE;
DROP TABLE IF EXISTS TRANSACTIONS CASCADE;
DROP TABLE IF EXISTS PRODUCTS CASCADE;
DROP TABLE IF EXISTS SELLER_BALANCES CASCADE;
DROP TABLE IF EXISTS FILE_METADATA CASCADE;
DROP TABLE IF EXISTS FILE_TRANSACTIONS CASCADE;

SELECT * FROM FILE_METADATA;
SELECT * FROM SELLERS;
SELECT * FROM TRANSACTIONS;
SELECT * FROM PRODUCTS;
SELECT * FROM SELLER_BALANCES;
SELECT * FROM FILE_TRANSACTIONS;
SELECT * FROM users;

SELECT s.id, s.name, sb.balance, sb.updated_at
FROM seller_balances sb
JOIN sellers s ON s.id = sb.seller_id
WHERE s.id = '81e53a87-472c-4993-bcef-b7aad8351724'

SELECT t.*
FROM transactions t
JOIN file_transactions ft ON t.id = ft.transaction_id
WHERE ft.file_id = '2895378b-ec54-4151-a9f5-bbefc1d662bf';

INSERT INTO users (id, name, email, password, role, created_at, updated_at)
VALUES (
    uuid_generate_v4(),
    'John Doe',
    'john.doe@example.com',
    hash_password('asdasdadsdas'),
    'user',
    NOW(),
    NOW()
);