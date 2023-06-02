ALTER TABLE public.c_e_file DROP CONSTRAINT file_rarity_fk,
DROP CONSTRAINT file_character_fk,
DROP CONSTRAINT file_equipment_fk;

COMMIT;