CREATE TABLE user_course_progress (
    user_id uuid NOT NULL,
    course_id uuid NOT NULL,
    resource_type varchar(255) NOT NULL,
    resource_id uuid NOT NULL,
    is_completed boolean NOT NULL,
    metadata json,
    created_at timestamp NOT NULL,
    updated_at timestamp,
    CONSTRAINT user_course_progress_pk PRIMARY KEY (user_id, course_id, resource_type, resource_id),
    CONSTRAINT user_course_progress_courses_fk FOREIGN KEY (course_id) REFERENCES public.courses(id) ON DELETE CASCADE ON UPDATE CASCADE
);