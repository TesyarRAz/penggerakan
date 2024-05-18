CREATE TABLE IF NOT EXISTS public.teachers (
	id uuid NOT NULL,
	user_id uuid NOT NULL,
	name varchar NULL,
	profile_image varchar NULL,
	created_at timestamp NOT NULL,
	updated_at timestamp,
	CONSTRAINT teachers_pk PRIMARY KEY (id),
	CONSTRAINT teachers_unique UNIQUE (user_id)
);