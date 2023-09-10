DROP TABLE IF EXISTS movies;
DROP TABLE IF EXISTS actors;
DROP TABLE IF EXISTS directors;
DROP TABLE IF EXISTS studio;
DROP TABLE IF EXISTS movie_directors;
DROP TABLE IF EXISTS movies_actors;

CREATE TYPE rating AS ENUM ('PG-10', 'PG-13', 'PG-18', 'R');
CREATE TABLE IF NOT EXISTS movies(
    id SERIAL PRIMARY KEY,
    title TEXT,
    year SMALLINT NOT NULL,
    box_office INT NOT NULL,
    studio_id INT REFERENCES studios(id),
    ratings rating
    UNIQUE(title, release_year)
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

CREATE TABLE movie_actors (
    movie_id INT REFERENCES movies(id),
    actor_id INT REFERENCES actors(id),
    PRIMARY KEY(movie_id, actor_id)
);

CREATE TABLE movie_directors (
    movie_id INT REFERENCES movies(id),
    director_id INT REFERENCES directors(id),
    PRIMARY KEY(movie_id, director_id)
);
----------------------------------------------------------
INSERT INTO movies(id, title,year,box_office,studio_id,ratings)
VALUES (1,'Мстители: Финал',2019,2799439100,1,'PG-13'), (2,'Мстители: Война бесконечности',2018,2052415039,1,'PG-13'), (3,'Джокер',2019,1074458282,2,'R');
----------------------------------------------------------
INSERT INTO actors(id,name,birth_date) VALUES(1,'Роберт Дауни-младший','04-04-1965'), (2,'Хоакин Феникс','28-10-1974'), (3,'Скарлет Йохансон', '1984-11-22');
----------------------------------------------------------
INSERT INTO directors(id,name,birth_date) VALUES(1,'Энтони Руссо','03-03-1970'), (2,'Тодд Филлипс','20-12-1970');
----------------------------------------------------------
INSERT INTO studio(id,title) VALUES(1,'Marvel Studios'), (2,'DC');
----------------------------------------------------------
INSERT INTO movie_actors (movie_id, actor_id) VALUES(1,1), (2,1), (3,2);
----------------------------------------------------------
INSERT INTO movies_directors(movie_id,director_id) VALUES(1,1),(2,1),(3,2);
----------------------------------------------------------
SELECT m.title, s.name AS studio_name
FROM movies m
         JOIN studios s ON m.studio_id = s.id;

SELECT a.name
FROM actors a
         JOIN movie_actors ma ON a.id = ma.actor_id
GROUP BY a.name
HAVING COUNT(ma.movie_id) > 2;

SELECT m.id AS movieId,m.title AS movieName,a.name AS actorName
FROM movies m INNER JOIN movies_actors_directors mad ON m.id=mad.movie_id
        INNER JOIN actors a ON mad.actor_id=a.id
    WHERE a.name='Роберт Дауни-младший';

SELECT * FROM movies m WHERE m.box_office>1000;
SELECT m.id,m.title,d.name
FROM movies m INNER JOIN movies_actors_directors mad ON m.id=mad.movie_id
            INNER JOIN directors d ON mad.director_id=d.id
        WHERE d.name='Энтони Руссо';