CREATE TABLE jobs (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    organization_id INTEGER NOT NULL REFERENCES organizations(id)
);

