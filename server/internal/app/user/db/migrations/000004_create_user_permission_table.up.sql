CREATE TABLE public.user_permission (
	user_id uuid NOT NULL,
	permission_id uuid NOT NULL,

	CONSTRAINT user_permission_pk PRIMARY KEY (user_id, permission_id)
);

ALTER TABLE public.user_permission ADD CONSTRAINT user_permission_permissions_fk FOREIGN KEY (permission_id) REFERENCES public.permissions(id) ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE public.user_permission ADD CONSTRAINT user_permission_users_fk FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE ON UPDATE CASCADE;