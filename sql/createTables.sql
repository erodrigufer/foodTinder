CREATE DATABASE ft;
\c ft
CREATE ROLE ft WITH LOGIN PASSWORD 'foodTinder';

CREATE TABLE IF NOT EXISTS sessions (
	id bigserial PRIMARY KEY,
   	session_id text NOT NULL
	);

CREATE TABLE IF NOT EXISTS products (
	id bigserial PRIMARY KEY,
	product_id text NOT NULL,
	product_name text NOT NULL
	);

CREATE TABLE IF NOT EXISTS votes (
	id bigserial PRIMARY KEY,
   	session_id text NOT NULL,
	product_id text NOT NULL,
	vote boolean
	);
