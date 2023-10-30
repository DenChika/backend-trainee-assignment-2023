CREATE TABLE users
(
    id serial primary key,
    username varchar(20) unique,
    password_hash varchar(255) unique
);

CREATE UNIQUE INDEX idx_users_username ON users(username, password_hash);