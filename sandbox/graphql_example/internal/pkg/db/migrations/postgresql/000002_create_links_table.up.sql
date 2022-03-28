CREATE TABLE IF NOT EXISTS links(
    id serial PRIMARY KEY,
    title VARCHAR (255),
    address VARCHAR (255),
    user_id serial,
    CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id)
);