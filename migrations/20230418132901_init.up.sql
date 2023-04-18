BEGIN;

-- Create DB
-- CREATE DATABASE go_orm;

-- Create Schema
CREATE SCHEMA go_orm;

-- Create Tables

CREATE TABLE go_orm.post(
	id UUID PRIMARY KEY,
	author TEXT NOT NULL,
	content TEXT NOT NULL,
	created_at TIMESTAMPTZ NOT NULL
);

CREATE TABLE go_orm.comment(
	id UUID PRIMARY KEY,
	author TEXT NOT NULL,
    content TEXT NOT NULL,
	post_id UUID REFERENCES go_orm.post (id),
	created_at TIMESTAMPTZ NOT NULL
);

COMMIT;