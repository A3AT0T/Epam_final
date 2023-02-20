CREATE TABLE user
(
    id       serial PRIMARY KEY,
    name     VARCHAR(30)         NOT NULL,
    surname  VARCHAR(40)         NOT NULL,
    email    VARCHAR(255) UNIQUE NOT NULL,
    pass     varchar(100)        NOT NULL,
    status   VARCHAR(6)          NOT NULL,
    is_admin BOOLEAN             NOT NULL DEFAULT FALSE
);

CREATE TABLE accounts
(
    id         serial PRIMARY KEY,
    acc        VARCHAR(31) UNIQUE NOT NULL,
    user_id    INT NO NULL,
    FOREIGN KEY (user_id) REFERENCES user (id),
    acc_status BOOLEAN DEFAULT FALSE,
    NOT NULL,
    amount     INT     DEFAULT 0
);

CREATE TABLE cards
(
    id      serial PRIMARY KEY,
    card_id VARCHAR(16) NOT NULL,
    acc_id  INT NO NULL,
    FOREIGN KEY (acc_id) REFERENCES accounts (id)

);

CREATE TABLE payments
(
    id     serial PRIMARY KEY,
    acc_id INT NO NULL,
    FOREIGN KEY (acc_id) REFERENCES accounts (id),
    amount INT NO NULL,
    status BOOLEAN   DEFAULT FALSE,
    date   TIMESTAMP DEFAULT NOW()
);

CREATE TABLE user_request
(
    id     serial PRIMARY KEY,
    acc_id INT NO NULL,
    FOREIGN KEY (acc_id) REFERENCES accounts (id),
    status BOOLEAN   DEFAULT FALSE,
    date   TIMESTAMP DEFAULT NOW()
);

CREATE TABLE log
(
    id      serial PRIMARY KEY,
    user_id INT NO NULL,
    FOREIGN KEY (user_id) REFERENCES user (id),
    massage VARCHAR(255),
    date    TIMESTAMP DEFAULT NOW()
)