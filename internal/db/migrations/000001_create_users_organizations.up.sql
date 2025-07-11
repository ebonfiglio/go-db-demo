CREATE TABLE organizations (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    organization_id BIGINT REFERENCES organizations(id) ON DELETE SET NULL,
);
