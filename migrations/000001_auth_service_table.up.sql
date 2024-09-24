CREATE TYPE attraction_type AS ENUM ('nature', 'park', 'beach', 'national parks', 'culture', 'museum', 'lake');

CREATE TABLE IF NOT EXISTS countries (
                                         id UUID DEFAULT gen_random_uuid(),
                                         country VARCHAR PRIMARY KEY,
                                         nationality VARCHAR,
                                         flag VARCHAR,
                                         deleted_at BIGINT DEFAULT 0 -- Added deleted_at for soft delete
);

CREATE TABLE IF NOT EXISTS history (
                                       id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                                       name VARCHAR(255) NOT NULL,
                                       description TEXT,
                                       country VARCHAR REFERENCES countries(country),
                                       image_url VARCHAR(255),
                                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                       deleted_at BIGINT DEFAULT 0 -- Added deleted_at for soft delete
);

CREATE TABLE IF NOT EXISTS attractions (
                                           id UUID DEFAULT gen_random_uuid(),
                                           category attraction_type,
                                           name VARCHAR(255) NOT NULL,
                                           description TEXT,
                                           country VARCHAR REFERENCES countries(country),
                                           location VARCHAR(255),
                                           image_url VARCHAR(255),
                                           created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                           updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                           deleted_at BIGINT DEFAULT 0 -- Added deleted_at for soft delete
);

CREATE TABLE IF NOT EXISTS foods (
                                     id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                                     food_name VARCHAR(255) NOT NULL UNIQUE, -- Ensure food names are unique
                                     food_type VARCHAR(100), -- Specify a max length for better control
                                     nationality VARCHAR(100),
                                     description TEXT,
                                     ingredients TEXT,
                                     image_url VARCHAR(255),
                                     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                     deleted_at BIGINT DEFAULT 0 -- Added deleted_at for soft delete
);
