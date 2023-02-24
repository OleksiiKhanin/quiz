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
    value VARCHAR(255) UNIQUE NOT NULL,
    description TEXT,
    type VARCHAR(50) NULL,
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

CREATE TABLE ratings (
    id BIGSERIAL PRIMARY KEY,
    card_id BIGSERIAL UNIQUE NOT NULL,
    stars REAL NOT NULL DEFAULT 0,
    shows INTEGER NOT NULL DEFAULT 0,
    modified_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_rating_card
        FOREIGN KEY(card_id)
            REFERENCES cards(id)
            ON DELETE CASCADE
);

CREATE INDEX ratings_card_id ON ratings USING btree (card_id);
