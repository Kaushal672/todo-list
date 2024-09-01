CREATE TYPE status AS ENUM ('not_started', 'in_progress', 'completed');

CREATE TABLE todos (
	todoId SERIAL PRIMARY KEY,
	title TEXT NOT NULL,
	currentStatus status DEFAULT 'not_started',
	userId INT NOT NULL REFERENCES users (userId) ON DELETE CASCADE,
	createdAt TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
	updatedAt TIMESTAMP WITH TIME ZONE
);