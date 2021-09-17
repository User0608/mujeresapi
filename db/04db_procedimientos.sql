
-- Procedimiento almacenado para verificar que
-- una cuanta de usuario, no esta siendo usada por
-- otra entidad o no existe.

create or replace function usuario_is_free(
	_usuario_id integer
) returns char(2) as $$
declare 
	_exit_user integer;
	_num_app_users integer;
	_num_instituciones integer;
	_num_colaborador integer;
	
BEGIN
	_exit_user = (select count(u) from usuario u where id=_usuario_id);
 	_num_app_users = (select count(ap) from app_user ap where id=_usuario_id);								
	_num_instituciones = (select count(i) from institucion i where usuario_id=_usuario_id);
	_num_colaborador = (select count(c) from colaborador c where usuario_id=_usuario_id);
	if _exit_user =0 then	
		return 'FK';
	elsif _num_app_users=0 and _num_instituciones=0 and _num_colaborador=0	then
		return 'OK';
	end	if;	
	return 'FF';
END;
$$ LANGUAGE plpgsql;

-- Funcion que permite varificar si una entidad existe o no.
create or replace function fn_consult(
	tablename varchar(30),
	table_id int
) returns integer as
$$
declare 
	resultado int; 
begin	
	execute 'select count(*) from ' ||
		quote_ident(tablename)||
		' where id = '||
		quote_nullable(table_id) into resultado;	
	return resultado;	
	EXCEPTION
		WHEN  undefined_table then
			raise exception 'La tabla % no existe!, funcion sql `fn_consult`',tablename;	
end;
$$ language plpgsql;

--- Funcion que me devuelve, la lista de notificaciones por id, y intitucion

create or replace function fn_find_notificacion_for_institucion(
	_notificacion_id int,
	_usuario_id int
) returns setof notificacion as
$$
begin
	return query select n.* from notificacion n 
			inner join institucion i on n.institucion_id =i.id
			inner join usuario u on i.usuario_id = u.id
				where n.id = _notificacion_id and u.id=_usuario_id;
end;
$$
LANGUAGE plpgsql;

-- Todas las notificaciones, correspondientes a un usuario institucional

create or replace function fn_find_all_notificacion_for_institucion(
	_usuario_id int
) returns setof notificacion as
$$
begin
	return query select n.* from notificacion n 
			inner join institucion i on n.institucion_id =i.id
			inner join usuario u on i.usuario_id = u.id
				where u.id=_usuario_id;
end;
$$
LANGUAGE plpgsql;
