-- migrate:up
ALTER TABLE public.c_e_file ADD CONSTRAINT file_rarity_fk FOREIGN KEY (rarity_id)
    REFERENCES public.rarity (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION,
ADD CONSTRAINT file_character_fk FOREIGN KEY (character_id)
    REFERENCES public.character (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION,
ADD CONSTRAINT file_equipment_fk FOREIGN KEY (equipment_id)
    REFERENCES public.equipment (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION;

ALTER TABLE public.card ADD c_e_file_id integer;

ALTER TABLE public.card ADD CONSTRAINT c_e_file_fk FOREIGN KEY (c_e_file_id)
    REFERENCES public.c_e_file (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION;

ALTER TABLE public.card DROP ipfs_uri;
ALTER TABLE public.card DROP image_uri;

COMMIT;