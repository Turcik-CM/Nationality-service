CREATE TYPE attraction_type AS ENUM ('nature', 'park', 'beach', 'national parks', 'culture', 'museum', 'lake');

CREATE TABLE IF NOT EXISTS countries
(
    id          UUID DEFAULT gen_random_uuid(),
    country     VARCHAR PRIMARY KEY,
    nationality VARCHAR,
    flag        VARCHAR
);

CREATE TABLE history
(
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name        VARCHAR(255) NOT NULL,
    description TEXT,
    country     VARCHAR REFERENCES countries(country),
    image_url   VARCHAR(255),
    created_at  TIMESTAMP        DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP        DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE attractions
(
    id          UUID      DEFAULT gen_random_uuid(),
    category    attraction_type,
    name        VARCHAR(255) NOT NULL,
    description TEXT,
    country     VARCHAR REFERENCES countries (country),
    location    VARCHAR(255),
    image_url   VARCHAR(255),
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
