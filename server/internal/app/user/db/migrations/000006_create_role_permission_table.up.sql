CREATE TABLE public.role_permission (
	role_id uuid NOT NULL,
	permission_id uuid NOT NULL,

	CONSTRAINT role_permission_pk PRIMARY KEY (role_id, permission_id)
);

ALTER TABLE public.role_permission ADD CONSTRAINT role_permission_permissions_fk FOREIGN KEY (permission_id) REFERENCES public.permissions(id) ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE public.role_permission ADD CONSTRAINT role_permission_roles_fk FOREIGN KEY (role_id) REFERENCES public.roles(id) ON DELETE CASCADE ON UPDATE CASCADE;