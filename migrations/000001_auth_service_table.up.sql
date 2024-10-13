CREATE TABLE IF NOT EXISTS attraction_types
(
    id       UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name     VARCHAR NOT NULL UNIQUE,
    activity INT
);


CREATE TABLE IF NOT EXISTS countries
(
    id   UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR NOT NULL UNIQUE,
    flag VARCHAR
);

CREATE TABLE IF NOT EXISTS cities
(
    id          UUID DEFAULT gen_random_uuid(),
    country_id  UUID REFERENCES countries (id) ON DELETE CASCADE,
    name        VARCHAR PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS history
(
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name        VARCHAR(255) NOT NULL,
    description TEXT,
    city        VARCHAR REFERENCES cities (name) ON DELETE CASCADE,
    image_url   VARCHAR(255),
    created_at  TIMESTAMP        DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP        DEFAULT CURRENT_TIMESTAMP,
    deleted_at  BIGINT           DEFAULT 0 -- Added deleted_at for soft delete
);

CREATE TABLE IF NOT EXISTS attractions
(
    id          UUID      DEFAULT gen_random_uuid(),
    category    VARCHAR REFERENCES attraction_types (name),
    name        VARCHAR(255) NOT NULL,
    description TEXT,
    city        VARCHAR REFERENCES cities (name) ON DELETE CASCADE,
    location    VARCHAR(255),
    image_url   VARCHAR(255),
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at  BIGINT    DEFAULT 0
);

CREATE TABLE IF NOT EXISTS foods
(
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    food_name   VARCHAR(255) NOT NULL UNIQUE,
    country_id  UUID REFERENCES countries (id) ON DELETE CASCADE,
    description TEXT,
    ingredients TEXT,
    image_url   VARCHAR(255),
    created_at  TIMESTAMP        DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP        DEFAULT CURRENT_TIMESTAMP,
    deleted_at  BIGINT           DEFAULT 0
);
