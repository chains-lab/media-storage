CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE media (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    folder TEXT NOT NULL,
    extension TEXT NOT NULL,
    resource_type TEXT NOT NULL,
    resource_id UUID NOT NULL,
    media_type TEXT NOT NULL,
    owner_id UUID,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);