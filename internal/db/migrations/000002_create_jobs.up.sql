CREATE TABLE jobs (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    organization_id BIGSERIAL NOT NULL REFERENCES organizations(id)
);

