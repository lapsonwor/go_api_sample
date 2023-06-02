
-- migrate:up
CREATE TABLE IF NOT EXISTS public.building
(
    id serial PRIMARY KEY,
    name character varying(255) COLLATE pg_catalog."default" NOT NULL,
    level integer NOT NULL DEFAULT 1,
    rarity character varying(255) COLLATE pg_catalog."default" NOT NULL,
    owner character varying(255) COLLATE pg_catalog."default" NOT NULL,
    isUR BOOLEAN NOT NULL,
    token_id integer NOT NULL,
    is_burned BOOLEAN NOT NULL DEFAULT false,
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone
);
