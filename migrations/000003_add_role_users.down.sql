BEGIN;

DROP TYPE roles;
ALTER TABLE groups DROP COLUMN role RESTRICT ;

COMMIT;