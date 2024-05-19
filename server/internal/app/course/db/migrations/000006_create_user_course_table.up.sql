CREATE TABLE public.user_course (
    user_id uuid NOT NULL,
    course_id uuid NOT NULL,
    is_completed boolean NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp,
    
    CONSTRAINT user_course_pk PRIMARY KEY (user_id, course_id)
);

ALTER TABLE public.user_course ADD CONSTRAINT user_course_courses_fk FOREIGN KEY (course_id) REFERENCES public.courses(id) ON DELETE CASCADE ON UPDATE CASCADE;