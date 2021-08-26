

create table usuario(
	id serial not null,
	username varchar(20) not null,
	password varchar(100) not null,
	estado 	boolean default true
);

create table roles(
	id serial not null,
	nombre varchar(80) not null,
	descripcion text
);

create table usuario_roles(
	usuario_id int not null,
	roles_id int not null	
);

alter table usuario 
	add constraint pk_usuario primary key(id);
	
alter table roles 
	add constraint pk_roles primary key(id);
	
alter table usuario_roles
	add constraint pk_usuario_roles primary key(usuario_id,roles_id),
	add constraint fk_usuario_roles__usuario foreign key(usuario_id)
		references usuario(id),
	add constraint fk_usuario_roles_roleas foreign key(roles_id)
		references roles(id);
		
-- datos base

insert into roles(nombre,descripcion) values
    ('admin','todos los permisos'),
	('user-t1','Usuario Aplicativo movil'),
	('user-t2','Usuario basico panel de control ');


insert into usuario(username,password)
	values('kevin002','maira002');	
	
insert into usuario_roles values (1,1);

-- Procedimiento almacenado para realizar login, el password no esta siendo encriptado
create or replace function ps_sign_in(
	_user_name varchar(60),
	_password varchar(60)
)
returns setof usuario
as
$$
	begin
	return query select  
		*		
	from usuario u 
	where u.username = _user_name
		and u.password =_password;
	end;
$$
language plpgsql;
