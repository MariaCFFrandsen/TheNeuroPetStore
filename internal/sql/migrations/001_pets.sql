CREATE TABLE pets (
    id        SERIAL PRIMARY KEY,
    name      VARCHAR NOT NULL,
    species   VARCHAR NOT NULL,
    imageurl  VARCHAR NOT NULL,
    qualities VARCHAR NOT NULL,
    available BOOLEAN NOT NULL
);