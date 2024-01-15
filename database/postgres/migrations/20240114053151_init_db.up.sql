CREATE TABLE IF NOT EXISTS app_user (
	app_user_id   SERIAL PRIMARY KEY,
	username      VARCHAR(255) NOT NULL,
	email         VARCHAR(255) NOT NULL,
	password_hash VARCHAR(255) NOT NULL,
	created_at    TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP);

CREATE TABLE IF NOT EXISTS project (
	project_id   SERIAL PRIMARY KEY,
	project_name VARCHAR(255) NOT NULL,
	description  TEXT,
	created_at   TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
	app_user_id  INT REFERENCES app_user (app_user_id) ON DELETE CASCADE);
