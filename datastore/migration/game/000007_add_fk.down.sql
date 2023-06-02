ALTER TABLE public.equipment DROP CONSTRAINT equip_type;

ALTER TABLE public.card DROP CONSTRAINT card_char_fk,
DROP CONSTRAINT card_equip_fk,
DROP CONSTRAINT card_rarity;

COMMIT;

