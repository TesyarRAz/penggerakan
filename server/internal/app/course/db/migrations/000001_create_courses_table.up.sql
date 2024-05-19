CREATE TABLE public.courses (
	id uuid NOT NULL,
	teacher_id uuid NOT NULL,
	image varchar NULL,
	"name" varchar NOT NULL,
	created_at timestamp NOT NULL,
	updated_at timestamp,
	CONSTRAINT courses_pk PRIMARY KEY (id)
);