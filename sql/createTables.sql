CREATE DATABASE ft;
\c ft
CREATE ROLE ft WITH LOGIN PASSWORD 'foodTinder';

CREATE TABLE IF NOT EXISTS sessions (
	id bigserial PRIMARY KEY,
   	session_id text NOT NULL
	);
