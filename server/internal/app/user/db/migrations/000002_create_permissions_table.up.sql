CREATE TABLE IF NOT EXISTS public.permissions (
	id uuid NOT NULL,
	"name" varchar NOT NULL,
	created_at timestamp NOT NULL,
	updated_at timestamp,
	CONSTRAINT permissions_pk PRIMARY KEY (id),
	CONSTRAINT permissions_unique UNIQUE (name)
);