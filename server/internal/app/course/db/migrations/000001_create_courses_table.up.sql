CREATE TABLE IF NOT EXISTS public.courses (
	id uuid NOT NULL,
	image varchar NULL,
	"name" varchar NOT NULL,
	created_at timestamp NOT NULL DEFAULT NOW(),
	updated_at timestamp,
	CONSTRAINT courses_pk PRIMARY KEY (id)
);