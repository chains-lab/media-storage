CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE media (
    ID            UUID PRIMARY KEY NOT NULL,
    format        TEXT NOT NULL,
    extensions     TEXT NOT NULL,
    size          INTEGER NOT NULL,
    url           TEXT NOT NULL,
    resource      TEXT NOT NULL,
    resource_id   TEXT NOT NULL,
    category      TEXT NOT NULL,
    owner_id      UUID NOT NULL ,
    created_at    TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE media_rules (
    ID            TEXt PRIMARY KEY NOT NULL,
    extensions    TEXT[]  NOT NULL,
    allowed_roles TEXT[]  NOT NULL,
    max_size      INTEGER NOT NULL,
    updated_at    TIMESTAMP NOT NULL DEFAULT now(),
    created_at    TIMESTAMP NOT NULL DEFAULT now()
);
