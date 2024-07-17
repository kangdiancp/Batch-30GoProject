-- if using postgre 10
--CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

--SELECT uuid_generate_v4();

--select gen_random_uuid ()

--drop table entity_type

create table entity_type(
	enty_id serial,
	entity_name varchar(25),
	constraint entity_id_pk primary key(enty_id),
	constraint entity_name_uq unique(entity_name)
)

insert into entity_type(entity_name)values ('BNI')

select * from entity_type

create table business_entity(
	buty_id serial,
	buty_name varchar(15),
	buty_enty_id integer,
	constraint buty_id_pk primary key(buty_id),
	constraint buty_enty_id_fk foreign key (buty_enty_id) references entity_type(enty_id)
)

insert into business_entity (buty_name,buty_enty_id) values('BNI',3)

select * from business_entity