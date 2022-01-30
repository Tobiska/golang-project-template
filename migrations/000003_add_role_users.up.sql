BEGIN;

CREATE TYPE roles AS ENUM ('Admin', 'Client');
ALTER TABLE users ADD COLUMN role roles;

COMMIT;