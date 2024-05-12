CREATE TABLE IF NOT EXISTS public.roles (
	id uuid NOT NULL,
	"name" varchar NULL,
	CONSTRAINT roles_pk PRIMARY KEY (id),
	CONSTRAINT roles_unique UNIQUE (name)
);