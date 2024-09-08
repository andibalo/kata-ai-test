DROP DATABASE IF EXISTS "pokemon-db-local";
CREATE DATABASE "pokemon-db-local";

\c "pokemon-db-local";


CREATE TABLE IF NOT EXISTS users  (
    id uuid primary key not null,
    name varchar(255) not null,
    email varchar(100) unique not null,
    last_accessed_at timestamptz not null,
    created_by varchar(100),
    created_at timestamptz not null,
    updated_by varchar(100),
    updated_at timestamptz,
    deleted_by varchar(100),
    deleted_at timestamptz
);