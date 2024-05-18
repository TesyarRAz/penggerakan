CREATE TABLE public.submodules (
	id uuid NOT NULL,
	module_id uuid NULL,
	"name" varchar NULL,
	"structure" varchar NULL,
    created_at timestamp NOT NULL DEFAULT NOW(),
    updated_at timestamp,
	CONSTRAINT submodules_pk PRIMARY KEY (id)
);

ALTER TABLE public.submodules ADD CONSTRAINT submodules_modules_fk FOREIGN KEY (module_id) REFERENCES public.modules(id);