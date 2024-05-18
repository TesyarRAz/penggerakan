CREATE TABLE public.content_pages (
	id uuid NOT NULL,
	course_id uuid NULL,
	title varchar NULL,
	"content" text NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp,
	CONSTRAINT content_blogs_pk PRIMARY KEY (id)
);

ALTER TABLE public.content_pages ADD CONSTRAINT content_pages_courses_fk FOREIGN KEY (course_id) REFERENCES public.courses(id);