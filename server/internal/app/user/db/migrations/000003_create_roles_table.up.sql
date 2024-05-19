CREATE TABLE public.roles (
	id uuid NOT NULL,
	"name" varchar NOT NULL,
	CONSTRAINT roles_pk PRIMARY KEY (id),
	CONSTRAINT roles_unique UNIQUE ("name")
);