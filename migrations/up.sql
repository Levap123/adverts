CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	email TEXT UNIQUE NOT NULL,
	password TEXT NOT NULL
);

-- ALTER TABLE users_workspaces ADD FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE;
-- ALTER TABLE users_workspaces ADD FOREIGN KEY (workspace_id) REFERENCES workspaces(id) ON DELETE CASCADE;
-- ALTER TABLE boards ADD FOREIGN KEY (workspace_id) REFERENCES workspaces(id) ON DELETE CASCADE;
-- ALTER TABLE lists ADD FOREIGN KEY (board_id) REFERENCES boards(id) ON DELETE CASCADE;
-- ALTER TABLE cards ADD FOREIGN KEY (list_id) REFERENCES lists(id) ON DELETE CASCADE;
