CREATE TABLE IF NOT EXISTS public.user_role (
	user_id uuid NULL,
	role_id uuid NULL
);

ALTER TABLE public.user_role ADD CONSTRAINT user_role_roles_fk FOREIGN KEY (role_id) REFERENCES public.roles(id) ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE public.user_role ADD CONSTRAINT user_role_users_fk FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE ON UPDATE CASCADE;