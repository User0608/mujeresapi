 create table direccion(
 	 id serial not null,
	 provincia  varchar(80) not null,
	 distrito varchar(80) not null,
	 direccion varchar(100) not null,
	 referencia varchar(450)
 );

create table app_user(
	 id int not null,
	 nombre varchar(80) not null,
	 apellido_paterno varchar(80) not null,
	 apellido_materno varchar(80) not null,
	 telefono varchar(20) default 'nn_nn',
	 dni char(8) default '00000000',
	 direccion_id int not null
 );
 
 create table alerta(
 	 id serial not null,
	 latitude varchar(20),
	 longitude varchar(20),
	 usuario_id int not null,
	 estado varchar(20) not null,
	 created_at timestamp not null,
	 updated_at timestamp,
	 deleted_at timestamp	 
 );	 
	
 create table multimedia(
 	id serial not null,
	alerta_id int not null,
	url varchar(100) not null,
	tipo varchar(6) not null	 
 ); 
 
alter table direccion
	add constraint pk_direccion primary key(id);
	
 alter table app_user 
 	add constraint pk_app_user primary key(id),
	add constraint fk_app_user__usuario
		foreign key(id) references usuario(id),
	add constraint fk_app_user__direccion 
		foreign key(direccion_id) references direccion(id);
-- --
alter table alerta
 	add constraint fk_alerta primary key(id),
	add constraint fk_alerta__app_user 
		foreign key(usuario_id) references app_user(id);
		
 alter table multimedia 
 	add constraint pk_multimedia primary key(id),
	add constraint fk_multimedia__alerta
		foreign key(alerta_id) references alerta(id);
