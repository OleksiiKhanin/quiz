CREATE TABLE images (
    id BIGSERIAL PRIMARY KEY,
    hash VARCHAR(60) UNIQUE NOT NULL,
    title VARCHAR(255) NOT NULL,
    type VARCHAR(50) NULL,
    data BYTEA NOT NULL,
    added_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX image_hash ON images USING hash (hash); 

CREATE TABLE cards (
    id BIGSERIAL PRIMARY KEY,
    value VARCHAR(255) NOT NULL,
    description TEXT,
    lang VARCHAR(80) NOT NULL,
    image_hash VARCHAR(60) DEFAULT NULL,
    added_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX cards_lang ON cards USING btree (lang);

CREATE TABLE pairs (
    id BIGSERIAL PRIMARY KEY,
    origin_id BIGSERIAL NOT NULL,
    pair_with BIGSERIAL NOT NULL,
    CONSTRAINT fk_origin_with
        FOREIGN KEY(origin_id)
            REFERENCES cards(id)
            ON DELETE CASCADE,
    CONSTRAINT fk_pair_with
        FOREIGN KEY(pair_with)
            REFERENCES cards(id)
            ON DELETE CASCADE
);

