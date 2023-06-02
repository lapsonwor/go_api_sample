-- migrate:up
CREATE TABLE IF NOT EXISTS public.c_e_file
(
    id serial PRIMARY KEY,
    rarity_id integer,
    character_id integer,
    equipment_id integer,
    ipfs_uri character varying(255) COLLATE pg_catalog."default" NOT NULL,
    image_uri character varying(255) COLLATE pg_catalog."default" NOT NULL,
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone
);