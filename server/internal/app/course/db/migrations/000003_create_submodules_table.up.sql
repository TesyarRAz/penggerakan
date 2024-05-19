CREATE TABLE public.submodules (
	id uuid NOT NULL,
	module_id uuid NOT NULL,
	"name" varchar NOT NULL,
	"structure" json NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp,
	CONSTRAINT submodules_pk PRIMARY KEY (id)
);

ALTER TABLE public.submodules ADD CONSTRAINT submodules_modules_fk FOREIGN KEY (module_id) REFERENCES public.modules(id);