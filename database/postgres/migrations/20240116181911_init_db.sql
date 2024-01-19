-- +goose Up
-- +goose StatementBegin

create table if not exists app_user (
	app_user_id   serial primary key,
	username      varchar(255) not null,
	email         varchar(255) not null,
	password_hash varchar(255) not null,
	created_at    TIMESTAMPTZ default current_timestamp);

create table if not exists project (
	project_id   serial primary key,
	project_name varchar(255) not null,
	description  text,
	created_at   TIMESTAMPTZ default current_timestamp,
	app_user_id  int references app_user (app_user_id) on delete cascade);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

drop table if exists project;
drop table if exists app_user;

-- +goose StatementEnd
