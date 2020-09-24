#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE DATABASE ecommerce;
EOSQL


psql -v ON_ERROR_STOP=1 --username postgres ecommerce <<-EOSQL
CREATE EXTENSION
IF NOT EXISTS "uuid-ossp";

CREATE TABLE users
(
   id UUID NOT NULL DEFAULT uuid_generate_v4(),
   birthday DATE NOT NULL,
   name VARCHAR NOT NULL,
   PRIMARY KEY(id)
);

INSERT INTO users
   (id, birthday, name)
VALUES
   ('8822888d-2713-4911-851b-c94b8aa60490', '1993-05-05', 'Roberta Sanches'),
   ('ff4404c7-f143-48ad-82c0-c85fd4bda3ae' , current_date, 'Renato Roberto');

DROP TABLE IF EXISTS products;
CREATE TABLE products
(
   id  SERIAL,
   sku UUID NOT NULL DEFAULT uuid_generate_v4(),
   name VARCHAR NOT NULL,
   price_in_cents integer,
   CONSTRAINT production UNIQUE(sku),
   PRIMARY KEY(id)
);

INSERT INTO products(name, price_in_cents)
VALUES ('Product 1', 1000),
('Product 2', 1000),
('Product 3', 1000),
('Product 4', 1000),
('Product 5', 1000),
('Product 6', 1000),
('Product 7', 1000),
('Product 8', 1000),
('Product 9', 1000),
('Product 10', 1000);
EOSQL