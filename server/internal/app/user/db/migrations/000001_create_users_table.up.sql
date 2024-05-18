CREATE TABLE IF NOT EXISTS public.users (
	id uuid NOT NULL,
	"name" varchar NOT NULL,
	email varchar NOT NULL,
	"password" varchar NOT NULL,
	profile_image varchar NULL,
	created_at timestamp NOT NULL DEFAULT NOW(),
	updated_at timestamp,
	CONSTRAINT users_pk PRIMARY KEY (id),
	CONSTRAINT users_unique UNIQUE (email)
);