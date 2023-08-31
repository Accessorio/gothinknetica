DROP TABLE IF EXISTS movies;
DROP TABLE IF EXISTS actors;
DROP TABLE IF EXISTS directors;
DROP TABLE IF EXISTS studio;
DROP TABLE IF EXISTS movies_studios;
DROP TABLE IF EXISTS movies_actors_directors;

CREATE TABLE IF NOT EXISTS movies_studios(
    movie_id INT NOT NULL,
    studio_id INT NOT NULL
);

CREATE UNIQUE INDEX movies_studios_unq ON movies_studios (movie_id,studio_id);
CREATE TABLE movies_actors_directors(
    movie_id INT NOT NULL,
    director_id INT NOT NULL,
    actor_id INT NOT NULL
);

CREATE UNIQUE INDEX movies_actors_directors_unq ON movies_actors_directors (
     movie_id,
     director_id,
     actor_id
    );
CREATE TYPE rating AS ENUM ('PG-10', 'PG-13', 'PG-18', 'R');
CREATE TABLE IF NOT EXISTS movies(
    id SERIAL PRIMARY KEY,
    title TEXT,
    year SMALLINT NOT NULL,
    box_office INT NOT NULL,
    ratings rating
);
CREATE TABLE IF NOT EXISTS actors(
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    birth_date DATE NOT NULL
);
CREATE TABLE IF NOT EXISTS directors(
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    birth_date DATE NOT NULL
);
CREATE TABLE IF NOT EXISTS studio(
    id SERIAL PRIMARY KEY,
    title TEXT
);
----------------------------------------------------------------------------
INSERT INTO movies(id, title,year,box_office,ratings)
VALUES (1,'Мстители: Финал',2019,2799439100,'PG-13');
INSERT INTO movies(id, title,year,box_office,ratings)
VALUES (2,'Мстители: Война бесконечности',2018,2052415039,'PG-13');
INSERT INTO movies(id, title,year,box_office,ratings)
VALUES (3,'Джокер',2019,1074458282,'R');
----------------------------------------------------------
INSERT INTO actors(id,name,birth_date) VALUES(1,'Роберт Дауни-младший','04-04-1965');
INSERT INTO actors(id,name,birth_date) VALUES(2,'Хоакин Феникс','28-10-1974');
----------------------------------------------------------
INSERT INTO directors(id,name,birth_date) VALUES(1,'Энтони Руссо','03-03-1970');
INSERT INTO directors(id,name,birth_date) VALUES(2,'Тодд Филлипс','20-12-1970');
----------------------------------------------------------
INSERT INTO studio(id,title) VALUES(1,'Marvel Studios');
INSERT INTO studio(id,title) VALUES(2,'DC');
----------------------------------------------------------
INSERT INTO movies_studios(movie_id,studio_id) VALUES(1,1);
INSERT INTO movies_studios(movie_id,studio_id) VALUES(2,1);
INSERT INTO movies_studios(movie_id,studio_id) VALUES(3,2);
----------------------------------------------------------
INSERT INTO movies_actors_directors(movie_id,director_id,actor_id) VALUES(1,1,1);
INSERT INTO movies_actors_directors(movie_id,director_id,actor_id) VALUES(2,1,1);
INSERT INTO movies_actors_directors(movie_id,director_id,actor_id) VALUES(3,2,2);
----------------------------------------------------------
SELECT m.id AS movieId,m.title AS movieName,a.name AS actorName
FROM movies m INNER JOIN movies_actors_directors mad ON m.id=mad.movie_id
        INNER JOIN actors a ON mad.actor_id=a.id
    WHERE a.name='Роберт Дауни-младший';
SELECT * FROM movies m WHERE m.box_office>1000;
SELECT m.id,m.title,d.name
FROM movies m INNER JOIN movies_actors_directors mad ON m.id=mad.movie_id
            INNER JOIN directors d ON mad.director_id=d.id
        WHERE d.name='Энтони Руссо';