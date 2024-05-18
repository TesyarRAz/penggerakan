CREATE TABLE IF NOT EXISTS public.roles (
	id uuid NOT NULL,
	"name" varchar NOT NULL,
	created_at timestamp NOT NULL,
	updated_at timestamp,
	CONSTRAINT roles_pk PRIMARY KEY (id),
	CONSTRAINT roles_unique UNIQUE (name)
);