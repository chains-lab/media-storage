CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE media (
    filename UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    extension TEXT NOT NULL,
    folder TEXT NOT NULL,
    resource_type TEXT NOT NULL,
    resource_id UUID NOT NULL,
    owner_id UUID NOT NULL ,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE media_rules (
    media_type    TEXT    PRIMARY KEY,
    max_size      BIGINT  NOT NULL,
    allowed_exits  TEXT[]  NOT NULL,
    folder   TEXT    NOT NULL,
    roles_access TEXT[] NOT NULL,
    updated_at    TIMESTAMP NOT NULL DEFAULT now(),
    created_at    TIMESTAMP NOT NULL DEFAULT now()
);
