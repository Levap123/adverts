CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	email TEXT UNIQUE NOT NULL,
	name TEXT UNIQUE NOT NULL,
	password TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS workspaces (
	id SERIAL PRIMARY KEY,
	title TEXT NOT NULL,
	logo TEXT
);

CREATE TABLE IF NOT EXISTS users_workspaces(
	user_id INT,
	workspace_id INT,	
	UNIQUE (user_id, workspace_id)
);

CREATE TABLE IF NOT EXISTS boards (
	id SERIAL PRIMARY KEY,
	title TEXT NOT NULL,
	background TEXT,
	workspace_id INT 
);

CREATE TABLE IF NOT EXISTS lists (
	id SERIAL PRIMARY KEY,
	title TEXT NOT NULL,
	board_id INT
);

CREATE TABLE IF NOT EXISTS cards (
	id SERIAL PRIMARY KEY,
	title TEXT NOT NULL,
	list_id INT
);

ALTER TABLE users_workspaces ADD FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE;
ALTER TABLE users_workspaces ADD FOREIGN KEY (workspace_id) REFERENCES workspaces(id) ON DELETE CASCADE;
ALTER TABLE boards ADD FOREIGN KEY (workspace_id) REFERENCES workspaces(id) ON DELETE CASCADE;
ALTER TABLE lists ADD FOREIGN KEY (board_id) REFERENCES boards(id) ON DELETE CASCADE;
ALTER TABLE cards ADD FOREIGN KEY (list_id) REFERENCES lists(id) ON DELETE CASCADE;
