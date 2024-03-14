CREATE TABLE IF NOT EXISTS "users" (
    "id"            serial PRIMARY KEY,
    "name"          varchar not null,
    "username"      varchar(255) not null unique,
    "password_hash" varchar(255) not null
);

CREATE TABLE IF NOT EXISTS "actors" (
    "id" serial PRIMARY KEY,
    "name" varchar(255),
    "sex" varchar(255) CHECK (sex = 'female' OR sex = 'male'),
    "birthday" date CHECK (birthday <= CURRENT_DATE),
    UNIQUE (name, sex, birthday)
);

CREATE TABLE IF NOT EXISTS "films" (
    "id" serial PRIMARY KEY,
    "name" varchar(255) CHECK (LENGTH(name) > 0 AND LENGTH(name) <= 150),
    "description" varchar(255) CHECK (LENGTH(description) <= 1000),
    "releaseDay" date CHECK (releaseDay <= CURRENT_DATE),
    "rating" int CHECK (rating >= 0 AND rating <= 10),
    UNIQUE (name, description, releaseDay, rating)
);

CREATE TABLE IF NOT EXISTS "actors_films" (
    id        serial PRIMARY KEY,
    actor_id int NOT NULL REFERENCES actors ON DELETE CASCADE,
    film_id   int NOT NULL REFERENCES films ON DELETE CASCADE,
    UNIQUE (actor_id, film_id)
);