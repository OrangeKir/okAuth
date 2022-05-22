CREATE TABLE users (
    login varchar(256) NOT NULL,
    password varchar(256),
    role roles,
    PRIMARY KEY(login)
)