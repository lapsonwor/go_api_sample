-- migrate:up
CREATE TABLE IF NOT EXISTS public.character
(
    id serial PRIMARY KEY,
    name character varying(255) COLLATE pg_catalog."default" NOT NULL,
    description json,
    attack_range integer NOT NULL DEFAULT 1,
    equipment_slot integer NOT NULL DEFAULT 1,
    -- is_ur bit(1) NOT NULL,
    -- tribe_id integer NOT NULL,
    country_id integer NOT NULL,
    gender_id integer NOT NULL,
    attribute_id integer NOT NULL,
    
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone


    -- CONSTRAINT character_tribe FOREIGN KEY (tribe_id)
    -- REFERENCES public.character_tribe (id) MATCH SIMPLE
    -- ON UPDATE NO ACTION
    -- ON DELETE NO ACTION,
    -- CONSTRAINT character_country FOREIGN KEY (country_id)
    -- REFERENCES public.character_country (id) MATCH SIMPLE
    -- ON UPDATE NO ACTION
    -- ON DELETE NO ACTION,
    -- CONSTRAINT character_gender FOREIGN KEY (gender_id)
    -- REFERENCES public.character_gender (id) MATCH SIMPLE
    -- ON UPDATE NO ACTION
    -- ON DELETE NO ACTION,
    -- CONSTRAINT character_attribute FOREIGN KEY (attribute_id)
    -- REFERENCES public.character_attribute (id) MATCH SIMPLE
    -- ON UPDATE NO ACTION
    -- ON DELETE NO ACTION,

);
