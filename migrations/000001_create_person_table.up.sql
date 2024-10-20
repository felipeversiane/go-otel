CREATE TABLE IF NOT EXISTS person (
    id UUID PRIMARY KEY NOT NULL,
    nickname VARCHAR(32) UNIQUE NOT NULL,
    name VARCHAR(100) NOT NULL,
    birthdate DATE NOT NULL,
    stack TEXT,
    search TEXT NOT NULL
);
