-- migrate:up
CREATE TABLE IF NOT EXISTS public.user
(
    id serial PRIMARY KEY,
    wallet_address character varying(255) COLLATE pg_catalog."default",
    email character varying(255) COLLATE pg_catalog."default" NOT NULL,
    name character varying(255) COLLATE pg_catalog."default",
    rank integer NOT NULL DEFAULT 1,
    balance integer NOT NULL DEFAULT 0,
    stamina integer NOT NULL DEFAULT 1,
    otp_auth_secret character varying(255) COLLATE pg_catalog."default",
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone
);
