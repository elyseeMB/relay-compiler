ALTER TABLE users
    ALTER COLUMN password TYPE BYTEA USING password::bytea;

ALTER TABLE users
    ALTER COLUMN password SET NOT NULL;
