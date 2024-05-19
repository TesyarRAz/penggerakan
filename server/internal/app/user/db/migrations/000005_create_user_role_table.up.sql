CREATE TABLE public.user_role (
	user_id uuid NOT NULL,
	role_id uuid NOT NULL,

	CONSTRAINT user_role_pk PRIMARY KEY (user_id, role_id)
);

ALTER TABLE public.user_role ADD CONSTRAINT user_role_roles_fk FOREIGN KEY (role_id) REFERENCES public.roles(id) ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE public.user_role ADD CONSTRAINT user_role_users_fk FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE ON UPDATE CASCADE;