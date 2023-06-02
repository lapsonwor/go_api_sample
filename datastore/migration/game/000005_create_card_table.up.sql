-- migrate:up
CREATE TABLE IF NOT EXISTS public.card
(
    id serial PRIMARY KEY,
    name character varying(255) COLLATE pg_catalog."default" NOT NULL,
    owner_address character varying(255) COLLATE pg_catalog."default" NOT NULL,
    equipment_id integer,
    character_id integer,
    level integer NOT NULL DEFAULT 1,
    token_id integer NOT NULL UNIQUE,
    rarity_id integer NOT NULL,
    ipfs_uri character varying(255) COLLATE pg_catalog."default" NOT NULL,
    image_uri character varying(255) COLLATE pg_catalog."default" NOT NULL,
    is_ur boolean NOT NULL DEFAULT FALSE,
    locked boolean  NOT NULL DEFAULT FALSE,
    number_of_skills integer NOT NULL DEFAULT 0,

    created_at timestamp without time zone NOT NULL DEFAULT now(),
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone

    -- CONSTRAINT equipment_level FOREIGN KEY (equipment_id, level)
    -- REFERENCES public.equipment_level (equipment_id, level) MATCH SIMPLE
    -- ON UPDATE NO ACTION
    -- ON DELETE NO ACTION,

);

