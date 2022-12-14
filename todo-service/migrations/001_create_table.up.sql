CREATE TABLE todos (
    id UUID PRIMARY KEY,
    user_id UUID REFERENCES users(id),
    body VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deadline TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    is_done BOOLEAN DEFAULT FALSE
);

DROP TABLE schema_migrations;