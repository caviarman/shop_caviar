-- +goose Up
-- +goose StatementBegin
create table if not exists "good"
(
	id          serial primary key,
	type        varchar(255),
	name        varchar(255),
	code        varchar(255),
	price       integer,
	status      varchar(255),
	created_at  timestamp,
	updated_at  timestamp
);

create table if not exists "container"
(
	id     serial primary key,
	name   varchar(255),
	code   varchar(255),
	weight integer
);

create table if not exists "order"
(
	id           serial primary key,
	user_id      integer,
	status       varchar(255),
	created_at   timestamp,
	updated_at   timestamp
);

create table if not exists "order_item"
(
	id           serial primary key,
	order_id     integer,
	good_id      integer,
	weight       integer,
	count        integer
);

create table if not exists "user"
(
	id                serial primary key,
	name              varchar(255),
	email             varchar(255),
	phone             varchar(255),
	telegram_username varchar(255),
	telegram_id       integer unique not null,
	created_at        timestamp,
	updated_at        timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists "user";
drop table if exists "order_item";
drop table if exists "order";
drop table if exists "container";
drop table if exists "good";
-- +goose StatementEnd
