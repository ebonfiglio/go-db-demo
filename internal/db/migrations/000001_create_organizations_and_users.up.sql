CREATE TABLE organizations (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

ALTER TABLE users
ADD COLUMN organization_id INTEGER REFERENCES organizations(id);