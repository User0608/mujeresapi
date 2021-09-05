create table institucion(
	id serial not null,
	persona varchar(120) not null,
	direccion_id int not null,
	telefono varchar(20) not null,
	email varchar(80) not null,
	tipo varchar(80) not null,
	usuario_id int not null
);

create table persona(
	id serial not null,
	nombre varchar(100) not null,
	apellido_materno varchar(100) not null,
	apellido_paterno varchar(100) not null,
	telefono varchar(20) not null,
	dni char(8) not null,
	direccion_id int not null
);

create table efectivo(
	id serial not null,
	institucion_id int not null,
	persona_id int not null
);

create table colaborador(
	id serial not null,
	persona_id int not null,
	usuario_id int not null
);
create table notificacion(
	id serial not null,
	alerta_id int not null,
	institucion_id int not null,
	titulo varchar(200) not null default 'none',
	nivel int not null default 2,
	descricion text,
	colaborador_id int not null
);

create table reporte(
	id serial not null,
	descripcion text not null,
	created_at timestamp not null,
	updated_at timestamp,
	deleted_at timestamp,
	notificacion_id int not null
);
create table asignacion(
	id serial not null,
	efectivo_id int not null,
	notificacion_id int not null,
	created_at timestamp not null,
	updated_at timestamp,
	deleted_at timestamp
);


alter table institucion
	add constraint pk_institucion primary key(id),
	add constraint fk_institucion__direccion
		foreign key(direccion_id) references direccion(id),
	add constraint fk_institucion_usuario
		foreign key(usuario_id) references usuario(id);
		
		
alter table persona
	add constraint pk_persona primary key(id),
	add constraint fk_persona__direccion
		foreign key(direccion_id) references direccion(id);

alter table efectivo
	add constraint pk_efectivo primary key(id),
	add constraint fk_efectivo__institucion
		foreign key(institucion_id) references institucion(id),
	add constraint fk_efectivo__persona 
		foreign key(persona_id) references persona(id);
		
alter table colaborador
	add constraint pk_colaborador primary key(id),
	add constraint fk_colaborador__persona
		foreign key(persona_id) references persona(id),
	add constraint fk_colaborador__usuario
		foreign key(usuario_id) references usuario(id);
		
alter table notificacion
	add constraint pk_notificacion primary key(id),
	add constraint fk_notification__alerta
		foreign key(alerta_id) references alerta(id),
	add constraint fk_notificacion__institucion
		foreign key(institucion_id) references institucion(id),
	add constraint fk_notification__colaborador
		foreign key(colaborador_id) references colaborador(id);

alter table reporte
	add constraint pk_reporte primary key(id),
	add constraint fk_reporte__notificacion 
		foreign key(notificacion_id) references notificacion(id);
		
alter table asignacion 
	add constraint pk_asignacion primary key(id),
	add constraint fk_asignacion__efectivo
		foreign key(efectivo_id) references efectivo(id),
	add constraint fk_asignacion__notificacion 
		foreign key(notificacion_id) references notificacion(id);