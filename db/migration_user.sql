-- creating user table in postgres DATABASE
CREATE TABLE users(
    id SERIAL NOT NULL PRIMARY KEY,
    last_name TEXT,
    first_name TEXT,
    email TEXT
);