CREATE TABLE IF NOT EXISTS attraction_type(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR PRIMARY KEY,
    activity int
);

CREATE TABLE IF NOT EXISTS countries
(
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    country     VARCHAR PRIMARY KEY,
    city        VARCHAR,
    nationality VARCHAR,
    flag        VARCHAR
);

CREATE TABLE IF NOT EXISTS history
(
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name        VARCHAR(255) NOT NULL,
    description TEXT,
    country     VARCHAR REFERENCES countries (country),
    image_url   VARCHAR(255),
    created_at  TIMESTAMP        DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP        DEFAULT CURRENT_TIMESTAMP,
    deleted_at  BIGINT           DEFAULT 0 -- Added deleted_at for soft delete
);

CREATE TABLE IF NOT EXISTS attractions
(
    id          UUID      DEFAULT gen_random_uuid(),
    category    VARCHAR REFERENCES attraction_type(name),
    name        VARCHAR(255) NOT NULL,
    description TEXT,
    country     VARCHAR REFERENCES countries (country),
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
    food_type   VARCHAR(100),
    nationality VARCHAR(100),
    description TEXT,
    ingredients TEXT,
    image_url   VARCHAR(255),
    created_at  TIMESTAMP        DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP        DEFAULT CURRENT_TIMESTAMP,
    deleted_at  BIGINT           DEFAULT 0
);
