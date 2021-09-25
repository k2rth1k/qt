CREATE TABLE IF NOT EXISTS accounts.users
(
    PRIMARY KEY (email, user_id),

    user_id                      VARCHAR(16)               NOT NULL,

    first_name                   VARCHAR(255)              NOT NULL,
    last_name                    VARCHAR(255)              NOT NULL,

    email                        VARCHAR(255)              NOT NULL,
    phone                        VARCHAR(255)              NOT NULL,
    password                     VARCHAR(255)              NOT NULL
);