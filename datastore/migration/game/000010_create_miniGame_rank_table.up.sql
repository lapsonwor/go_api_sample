CREATE TABLE IF NOT EXISTS public.miniGame_rank
(
    id serial PRIMARY KEY,
    wallet_address character varying(255) COLLATE pg_catalog."default" UNIQUE NOT NULL,
    mark integer NOT NULL DEFAULT 0,
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone
);

