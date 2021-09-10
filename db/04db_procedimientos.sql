
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