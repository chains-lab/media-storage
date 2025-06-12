-- +migrate Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE media_rules (
    resource       TEXT      NOT NULL,
    category       TEXT      NOT NULL,
    created_at     TIMESTAMP NOT NULL DEFAULT now(),
    PRIMARY KEY (resource, category)
);

CREATE TABLE allowed_extensions (
    resource    TEXT      NOT NULL,
    category    TEXT      NOT NULL,
    extension   TEXT      NOT NULL,
    max_size    INTEGER   NOT NULL DEFAULT 10485760, -- 10 MB
    created_at  TIMESTAMP NOT NULL DEFAULT now(),
    updated_at  TIMESTAMP NOT NULL DEFAULT now(),

    PRIMARY KEY (resource, category, extension),
    FOREIGN KEY (resource, category)
        REFERENCES media_rules (resource, category)
        ON DELETE CASCADE
);

CREATE INDEX idx_allowed_ext_extension
    ON allowed_extensions (extension);

-- +migrate Down
DROP INDEX IF EXISTS idx_allowed_ext_extension;

DROP TABLE IF EXISTS allowed_extensions CASCADE;
DROP TABLE IF EXISTS media_rules CASCADE;

DROP EXTENSION IF EXISTS "uuid-ossp";
