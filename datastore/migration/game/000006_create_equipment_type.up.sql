-- migrate:up
CREATE TABLE IF NOT EXISTS public.equipment_type
(
    id serial PRIMARY KEY,
    name character varying(255) COLLATE pg_catalog."default" NOT NULL,
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone
);

