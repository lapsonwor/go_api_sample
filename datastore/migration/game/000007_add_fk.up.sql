
ALTER TABLE public.equipment ADD CONSTRAINT equip_type FOREIGN KEY (type_id)
    REFERENCES public.equipment_type (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION;

ALTER TABLE public.card ADD CONSTRAINT card_rarity FOREIGN KEY (rarity_id)
    REFERENCES public.rarity (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION,

ADD CONSTRAINT card_char_fk FOREIGN KEY (character_id)
    REFERENCES public.character (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION,

ADD CONSTRAINT card_equip_fk FOREIGN KEY (equipment_id)
    REFERENCES public.equipment (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION;

COMMIT;