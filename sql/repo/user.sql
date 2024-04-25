CREATE TABLE users (
    uuid varchar(255) NOT NULL UNIQUE,
    username varchar(255) NOT NULL UNIQUE,
    email varchar(255) NOT NULL UNIQUE,
    PRIMARY KEY (uuid)
);