create database users;


create table users(
	id serial primary key,
	first_name varchar,
	last_name varchar,
	gender varchar,
	age int,
	contact varchar,
	email varchar,
	password varchar,
	is_member bool default false,
	priority int default 1,
	verified bool default true,
	created_at timestamp,
	updated_at timestamp,
    deleted_at timestamp
)

create table post_Types (
	post_Type_id serial primary key,
	post_Type_Title text,
	created_at timestamp,
	updated_at timestamp,
    deleted_at timestamp
)

create table public.posts (
	post_id serial primary key,
	user_id int not null,
	post_name text,
	post_description text,
	post_type text[],
	constraint fk_users
	foreign key(user_id) references users(user_id)
)


