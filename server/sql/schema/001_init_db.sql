
CREATE TABLE public.genres (
    id integer NOT NULL PRIMARY KEY,
    genre character varying(255),
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);

CREATE TABLE public.movies (
    id integer NOT NULL PRIMARY KEY,
    title character varying(512),
    release_date date,
    runtime integer,
    mpaa_rating character varying(10),
    description text,
    image character varying(255),
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);

CREATE TABLE public.movies_genres (
    id integer NOT NULL PRIMARY KEY,
    movie_id integer NOT NULL,
    genre_id integer NOT NULL,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    CONSTRAINT fk_movie FOREIGN KEY (movie_id) REFERENCES public.movies (id) ON UPDATE CASCADE ON DELETE CASCADE,
    CONSTRAINT fk_genre FOREIGN KEY (genre_id) REFERENCES public.genres (id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE public.users (
    id integer NOT NULL PRIMARY KEY,
    first_name character varying(255),
    last_name character varying(255),
    email character varying(255),
    password character varying(255),
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);
