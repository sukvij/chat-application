create table if not exists users(
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
	image varchar
);

CREATE TABLE friends (
    user_id PRIMARY KEY,,
    friends_list JSONB NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

CREATE TABLE friends (
    id serial PRIMARY KEY,
	from int,
	to int,
	detail varchar
);


