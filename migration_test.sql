
DROP DATABASE IF EXISTS bondstate_db_test;
CREATE DATABASE bondstate_db_test;

\c bondstate_db_test;

DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS portfolios;
DROP TABLE IF EXISTS entries;

CREATE TABLE users (
    id INT GENERATED ALWAYS AS IDENTITY,
    email VARCHAR(50) NOT NULL UNIQUE,
    name VARCHAR(50),
    verified boolean,
    locked boolean,
    password VARCHAR(255) NOT NULL,
    PRIMARY KEY(id)
);

CREATE TABLE portfolios (
    id INT GENERATED ALWAYS AS IDENTITY,
    user_id INT UNIQUE,
    name VARCHAR(50) NOT NULL,
    PRIMARY KEY(id),
    CONSTRAINT fk_user_id FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE entries (
    id INT GENERATED ALWAYS AS IDENTITY,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    folio_id INTEGER NOT NULL,
    coin_name VARCHAR(50) NOT NULL,
    amount INTEGER NOT NULL,
    price DECIMAL NOT NULL,
    transaction_fee DECIMAL NOT NULL,
    PRIMARY KEY(id),
    CONSTRAINT fk_folio_id FOREIGN KEY(folio_id) REFERENCES portfolios(id) ON DELETE CASCADE
);

INSERT INTO users (email, name, verified, locked, password) values 
('vikas@bondstate.com', 'vikas', true, true, 'SUPERSECRETPASSWORD'),
('sella@bondstate.com', 'sella', true, true, 'SUPERSECRETPASSWORD'),
('newuser@bondstate.com', 'newuser', true, true, 'SUPERSECRETPASSWORD'),
('raj@bondstate.com', 'raj', true, true, 'SUPERSECRETPASSWORD');

INSERT INTO portfolios (user_id, name) values
(1, 'v1p'),
(2, 's1p'),
(3, 'n1p'),
(4, 'r1p');

insert into entries (folio_id, coin_name, amount, price , transaction_fee ) values
(1, 'ada', 100, 1.43, 10),
(2, 'ada', 100, 1.43, 10),
(3, 'ada', 100, 1.43, 10),
(4, 'ada', 100, 1.43, 10);