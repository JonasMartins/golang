CREATE TABLE IF NOT EXISTS public.contests (
    id SERIAL PRIMARY KEY,
    realese_date timestamp without time zone,
    bola_1 integer NOT NULL,
    bola_2 integer NOT NULL,
    bola_3 integer NOT NULL,
    bola_4 integer NOT NULL,
    bola_5 integer NOT NULL,
    bola_6 integer NOT NULL
);