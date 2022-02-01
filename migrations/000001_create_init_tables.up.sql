BEGIN;

CREATE TABLE groups (
  uuid uuid primary key,
  name VARCHAR(100) not null
);

CREATE TABLE users (
    id serial not null unique primary key,
    name varchar(100) not null,
    email varchar(100) not null unique,
    password_encrypted varchar(256) not null,
    group_id uuid null,
    CONSTRAINT group_fk FOREIGN KEY (group_id) REFERENCES groups(uuid)
);

CREATE TABLE sessions (
    uuid uuid primary key,
    group_id uuid null,
    CONSTRAINT group_fk FOREIGN KEY (group_id) REFERENCES groups(uuid)
);
COMMIT;