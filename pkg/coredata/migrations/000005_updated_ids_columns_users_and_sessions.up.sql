ALTER TABLE sessions DROP CONSTRAINT sessions_user_id_fkey;

ALTER TABLE users 
    ALTER COLUMN id TYPE TEXT USING id::text;

ALTER TABLE sessions 
    ALTER COLUMN id TYPE TEXT USING id::text,
    ALTER COLUMN user_id TYPE TEXT USING user_id::text;

ALTER TABLE sessions
    ADD CONSTRAINT sessions_user_id_fkey
    FOREIGN KEY (user_id) REFERENCES users(id);
