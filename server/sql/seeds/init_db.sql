-- Initial Seed Data
-- Seed data for genres table
INSERT INTO public.genres (id, genre, created_at, updated_at) VALUES
(1, 'Comedy', '2022-09-23 00:00:00', '2022-09-23 00:00:00'),
(2, 'Sci-Fi', '2022-09-23 00:00:00', '2022-09-23 00:00:00'),
(3, 'Horror', '2022-09-23 00:00:00', '2022-09-23 00:00:00'),
(4, 'Romance', '2022-09-23 00:00:00', '2022-09-23 00:00:00'),
(5, 'Action', '2022-09-23 00:00:00', '2022-09-23 00:00:00'),
(6, 'Thriller', '2022-09-23 00:00:00', '2022-09-23 00:00:00'),
(7, 'Drama', '2022-09-23 00:00:00', '2022-09-23 00:00:00'),
(8, 'Mystery', '2022-09-23 00:00:00', '2022-09-23 00:00:00'),
(9, 'Crime', '2022-09-23 00:00:00', '2022-09-23 00:00:00'),
(10, 'Animation', '2022-09-23 00:00:00', '2022-09-23 00:00:00'),
(11, 'Adventure', '2022-09-23 00:00:00', '2022-09-23 00:00:00'),
(12, 'Fantasy', '2022-09-23 00:00:00', '2022-09-23 00:00:00'),
(13, 'Superhero', '2022-09-23 00:00:00', '2022-09-23 00:00:00');

-- Seed data for movies table
INSERT INTO public.movies (id, title, release_date, runtime, mpaa_rating, description, image, created_at, updated_at) VALUES
(1, 'Highlander', '1986-03-07', 116, 'R', 'He fought his first battle on the Scottish Highlands in 1536. He will fight his greatest battle on the streets of New York City in 1986. His name is Connor MacLeod. He is immortal.', '/8Z8dptJEypuLoOQro1WugD855YE.jpg', '2022-09-23 00:00:00', '2022-09-23 00:00:00'),
(2, 'Raiders of the Lost Ark', '1981-06-12', 115, 'PG-13', 'Archaeology professor Indiana Jones ventures to seize a biblical artefact known as the Ark of the Covenant. While doing so, he puts up a fight against Renee and a troop of Nazis.', '/ceG9VzoRAVGwivFU403Wc3AHRys.jpg', '2022-09-23 00:00:00', '2022-09-23 00:00:00'),
(3, 'The Godfather', '1972-03-24', 175, '18A', 'The aging patriarch of an organized crime dynasty in postwar New York City transfers control of his clandestine empire to his reluctant youngest son.', '/3bhkrj58Vtu7enYsRolD1fZdja1.jpg', '2022-09-23 00:00:00', '2022-09-23 00:00:00');

-- Seed data for movies_genres table
INSERT INTO public.movies_genres (id, movie_id, genre_id, created_at, updated_at) VALUES
(1, 1, 5, '2022-09-23 00:00:00', '2022-09-23 00:00:00'),
(2, 1, 12, '2022-09-23 00:00:00', '2022-09-23 00:00:00'),
(3, 2, 5, '2022-09-23 00:00:00', '2022-09-23 00:00:00'),
(4, 2, 11, '2022-09-23 00:00:00', '2022-09-23 00:00:00'),
(5, 3, 9, '2022-09-23 00:00:00', '2022-09-23 00:00:00'),
(6, 3, 7, '2022-09-23 00:00:00', '2022-09-23 00:00:00');

-- Seed data for users table
INSERT INTO public.users (id, first_name, last_name, email, password, created_at, updated_at) VALUES
(1, 'Admin', 'User', 'admin@example.com', '$2a$14$wVsaPvJnJJsomWArouWCtusem6S/.Gauq/GjOIEHpyh2DAMmso1wy', '2022-09-23 00:00:00', '2022-09-23 00:00:00');

