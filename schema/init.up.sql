CREATE TABLE users
(
    id serial NOT NULL UNIQUE,
    username VARCHAR(255),
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL
);