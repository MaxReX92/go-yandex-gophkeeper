CREATE TABLE IF NOT EXISTS secret (
	id TEXT NOT NULL PRIMARY KEY,
	userId INTEGER NOT NULL,
	content BYTEA NOT NULL
);
