CREATE TABLE IF NOT EXISTS public.permissions (
	id uuid NOT NULL,
	"name" varchar NULL,
	CONSTRAINT permissions_pk PRIMARY KEY (id),
	CONSTRAINT permissions_unique UNIQUE (name)
);