CREATE TABLE public.content_videos (
	id uuid NOT NULL,
	course_id uuid NULL,
	video_url varchar NULL,
    title varchar NULL,
    description text NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp,
	CONSTRAINT content_video_pk PRIMARY KEY (id)
);

ALTER TABLE public.content_videos ADD CONSTRAINT content_video_courses_fk FOREIGN KEY (course_id) REFERENCES public.courses(id);