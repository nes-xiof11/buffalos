-- +migrate up

CREATE TABLE IF NOT EXISTS projects
(
    id         BIGSERIAL PRIMARY KEY,
    owner_id   BIGINT       NOT NULL,
    name       VARCHAR(255) NOT NULL,
    created_at TIMESTAMP    NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP    NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_projects_owner
        FOREIGN KEY (owner_id) REFERENCES users (id)
            ON DELETE CASCADE
);

CREATE INDEX idx_projects_owner_id ON projects (owner_id);
CREATE INDEX idx_projects_name ON projects (name);
CREATE INDEX idx_projects_created_at ON projects (created_at);


-- +migrate down