package test

const merchantTableCreationQuery = `CREATE TABLE IF NOT EXISTS merchant
(
	id SERIAL PRIMARY KEY,
	code text NOT NULL,
	name text NOT NULL,
	CONSTRAINT merchant_code_unique UNIQUE ("code")
)`

const teamTableCreationQuery = `CREATE TABLE IF NOT EXISTS team
(
	id SERIAL PRIMARY KEY,
	email text NOT NULL,
	merchant_id INT NULL,
	CONSTRAINT merchant_code_unique UNIQUE ("code")
)`

const APP_DB_USERNAME = "postgres"
const APP_DB_NAME = "postgres"
const APP_DB_HOST = "database"
const APP_DB_PASSWORD = "postgres"
const APP_DB_PORT = 5432