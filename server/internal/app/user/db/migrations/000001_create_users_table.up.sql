CREATE TABLE IF NOT EXISTS public.users (
	id uuid NOT NULL,
	"name" varchar NULL,
	email varchar NULL,
	"password" varchar NULL,
	profile_image varchar NULL,
	CONSTRAINT users_pk PRIMARY KEY (id),
	CONSTRAINT users_unique UNIQUE (email)
);