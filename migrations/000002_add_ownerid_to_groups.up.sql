BEGIN;

ALTER TABLE groups ADD COLUMN owner_id int;
ALTER TABLE groups ADD CONSTRAINT user_fk FOREIGN KEY (owner_id) REFERENCES users (id);

COMMIT;