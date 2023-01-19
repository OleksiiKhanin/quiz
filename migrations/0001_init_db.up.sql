CREATE TABLE images (
    id BIGSERIAL PRIMARY KEY,
    tittle VARCHAR(255) NOT NULL,
    hash VARCHAR(60) UNIQUE NOT NULL,
    data BYTEA NOT NULL,
    added_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE languages (
    id BIGSERIAL PRIMARY KEY,
    value VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    lang VARCHAR(100) NOT NULL,
    image_uuid VARCHAR(60) DEFAULT NULL,
    added_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX languages_lang ON languages USING btree (lang);

CREATE TABLE compliances (
    id BIGSERIAL PRIMARY KEY,
    origin_id SERIAL NOT NULL,
    compliance_with SERIAL NOT NULL
);