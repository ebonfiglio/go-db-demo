CREATE TABLE organizations (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    job_id BIGSERIAL REFERENCES jobs(id) ON DELETE CASCADE,
    organization_id BIGSERIAL REFERENCES organizations(id) ON DELETE SET NULL,
);
