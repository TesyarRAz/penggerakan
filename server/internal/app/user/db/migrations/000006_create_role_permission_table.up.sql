CREATE TABLE IF NOT EXISTS public.role_permission (
	role_id uuid NULL,
	permission_id uuid NULL
);

ALTER TABLE public.role_permission ADD CONSTRAINT role_permission_permissions_fk FOREIGN KEY (permission_id) REFERENCES public.permissions(id) ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE public.role_permission ADD CONSTRAINT role_permission_roles_fk FOREIGN KEY (role_id) REFERENCES public.roles(id) ON DELETE CASCADE ON UPDATE CASCADE;