CREATE TABLE public.modules (
	id uuid NOT NULL,
	course_id uuid NOT NULL,
	"name" varchar NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp,
	CONSTRAINT modules_pk PRIMARY KEY (id)
);

ALTER TABLE public.modules ADD CONSTRAINT modules_courses_fk FOREIGN KEY (course_id) REFERENCES public.courses(id);