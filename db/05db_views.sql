-- Devuelve los usuarios, no asignados con ninguna entidad.
create or replace view vw_usuarios_libres 
	as select u.* from usuario u
		left join app_user au
		on u.id = au.id
		left join colaborador c
		on u.id = c.usuario_id
		left join institucion i
		on u.id = i.usuario_id
		where au.id is null and 
			c.usuario_id is null and 
			i.usuario_id is null

select * from vw_usuarios_libres
	