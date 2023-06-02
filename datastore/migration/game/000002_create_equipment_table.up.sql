-- migrate:up
CREATE TABLE IF NOT EXISTS public.equipment
(
    id serial PRIMARY KEY,
    name character varying(255) COLLATE pg_catalog."default" NOT NULL,
    critical_rate integer NOT NULL DEFAULT 0,
    attack_range integer NOT NULL DEFAULT 1,
    description json,
    type_id integer NOT NULL,

    created_at timestamp without time zone NOT NULL DEFAULT now(),
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone

);

