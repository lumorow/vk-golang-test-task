CREATE TABLE IF NOT EXISTS "users" (
    "id"            serial PRIMARY KEY,
    "username"      varchar(255) not null unique,
    "password_hash" varchar(255) not null,
    "role" VARCHAR(20) NOT NULL
);

CREATE TABLE IF NOT EXISTS "actors" (
    "id" serial PRIMARY KEY,
    "name" varchar(255),
    "sex" varchar(255),
    "birthday" date,
    UNIQUE (name, sex, birthday)
);

CREATE TABLE IF NOT EXISTS "films" (
    "id" serial PRIMARY KEY,
    "name" varchar(255),
    "description" varchar(255),
    "release" date,
    "rating" int,
    UNIQUE (name, description, release, rating)
);

CREATE TABLE IF NOT EXISTS "actors_films" (
    id        serial PRIMARY KEY,
    actor_id int NOT NULL REFERENCES actors ON DELETE CASCADE,
    film_id   int NOT NULL REFERENCES films ON DELETE CASCADE,
    UNIQUE (actor_id, film_id)
);