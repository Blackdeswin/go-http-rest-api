CREATE TABLE users (
    id bigserial NOT NULL PRIMARY KEY,
    email VARCHAR not NULL UNIQUE,
    encrypted_password VARCHAR NOT NULL
);