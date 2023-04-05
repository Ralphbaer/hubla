SET search_path = public, pgcrypto;

CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Table to store the types of transactions
CREATE TYPE transaction_type AS ENUM (
    'CREATOR_SALE',
    'AFFILIATE_SALE',
    'COMMISSION_PAID',
    'COMMISSION_RECEIVED'
);

CREATE TYPE seller_type AS ENUM ('CREATOR', 'AFFILIATE');

-- Table to store file metadata
CREATE TABLE IF NOT EXISTS file_metadata (
    id UUID PRIMARY KEY,
    file_size BIGINT NOT NULL,
    disposition TEXT NOT NULL,
    hash TEXT NOT NULL UNIQUE,
    binary_data BYTEA NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);

-- Table to store seller
CREATE TABLE IF NOT EXISTS seller (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    seller_type seller_type NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);

-- Table to store product
CREATE TABLE IF NOT EXISTS product (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    creator_id UUID NOT NULL REFERENCES seller(id),
    created_at TIMESTAMP NOT NULL DEFAULT now()
);

-- Table to store transactions
CREATE TABLE IF NOT EXISTS transaction_record (
    id UUID PRIMARY KEY,
    t_type transaction_type NOT NULL,
    t_date TIMESTAMP NOT NULL,
    product_id UUID NOT NULL REFERENCES product(id),
    amount NUMERIC(19, 2) NOT NULL,
    seller_id UUID NOT NULL REFERENCES seller(id),
    created_at TIMESTAMP NOT NULL DEFAULT now()
);

-- Table that relates file_id with transaction_record
CREATE TABLE IF NOT EXISTS file_transaction (
    id UUID PRIMARY KEY,
    file_id UUID NOT NULL,
    transaction_id UUID NOT NULL,
    FOREIGN KEY (file_id) REFERENCES file_metadata(id),
    FOREIGN KEY (transaction_id) REFERENCES transaction_record(id),
    UNIQUE (file_id, transaction_id)
);

-- Table to store the seller balances
CREATE TABLE IF NOT EXISTS seller_balance (
    id UUID PRIMARY KEY,
    seller_id UUID NOT NULL REFERENCES seller(id),
    balance NUMERIC(19, 2) NOT NULL,
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    UNIQUE(seller_id)
);

-- Table to store users
CREATE TABLE IF NOT EXISTS user_account (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- Indexes for better performance
CREATE INDEX IF NOT EXISTS idx_file_metadata_hash ON file_metadata(hash);

CREATE INDEX IF NOT EXISTS idx_product_name ON product(name);

CREATE INDEX IF NOT EXISTS idx_seller_name ON seller(name);

CREATE INDEX IF NOT EXISTS idx_user_account_email ON user_account(email);

INSERT INTO user_account (id, name, email, password, role, created_at, updated_at)
VALUES (
    uuid_generate_v4(),
    'Marie Curie',
    'marie.curie@hub.la',
    PGP_SYM_ENCRYPT('radiantforce', 'AES_KEY'),
    'admin',
    NOW(),
    NOW()
);

INSERT INTO user_account (id, name, email, password, role, created_at, updated_at)
VALUES (
    uuid_generate_v4(),
    'Nikola Tesla',
    'nikola.testa@hub.la',
    PGP_SYM_ENCRYPT('radiantforce', 'AES_KEY'),
    'admin',
    NOW(),
    NOW()
);

INSERT INTO user_account (id, name, email, password, role, created_at, updated_at)
VALUES (
    uuid_generate_v4(),
    'Rosalind Elsie Franklin',
    'rosalind.franklin@hub.la',
    PGP_SYM_ENCRYPT('helixstructure1953', 'AES_KEY'),
    'admin',
    NOW(),
    NOW()
);