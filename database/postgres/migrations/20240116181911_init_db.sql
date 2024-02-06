-- +goose Up
-- +goose StatementBegin

create table if not exists products_order (
    id             serial,
    weight_payment money not null,
    taxes          money not null,
    total_amount   money not null,
    arrived_at     timestamptz default current_timestamp,
    created_at     timestamptz default current_timestamp,
    updated_at     timestamptz default current_timestamp,
    primary key (id));

create table if not exists product_category (
    id          serial,
    name        varchar(255) not null,
    description varchar(255) not null,
    primary key (id));

create table if not exists product_subcategory (
    id                 serial,
    parent_category_id serial       not null,
    name               varchar(255) not null,
    description        varchar(255) not null,
    primary key (id),
    constraint fk_parent_category foreign key (parent_category_id) references product_category (id));

create table if not exists product_size (
    code        varchar(5),
    description varchar(255) not null,
    primary key (code));

create table if not exists product_color (
    id   serial,
    name varchar(255) not null,
    hex  varchar(10)  not null,
    primary key (id));

create table if not exists product_source (
    id   serial,
    name varchar(255) not null,
    primary key (id));

create table if not exists product (
    id             serial,
    subcategory_id serial       not null,
    name           varchar(255) not null,
    description    varchar(255),
    price          money        not null,
    cost           money        not null,
    quantity       int          not null,
    size_code      char         not null,
    color_id       serial       not null,
    brand          varchar(255) not null,
    sku            varchar(12)  not null,
    upc            varchar(12)  not null,
    image_url      varchar(255) not null,
    source_id      serial       not null,
    source_url     varchar(300),
    is_offered     boolean      not null default false,
    offer_percent  int                   default 0,
    is_active      boolean      not null default true,
    created_at     timestamptz           default current_timestamp,
    updated_at     timestamptz           default current_timestamp,
    primary key (id),
    constraint fk_category foreign key (subcategory_id) references product_subcategory (id),
    constraint fk_size foreign key (size_code) references product_size (code),
    constraint fk_color foreign key (color_id) references product_color (id),
    constraint fk_source foreign key (source_id) references product_source (id));

create table product_in_order (
    order_id   serial not null,
    product_id serial not null,
    constraint fk_order foreign key (order_id) references products_order (id),
    constraint fk_product foreign key (product_id) references product (id));

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

drop table if exists product_in_order;
drop table if exists product;
drop table if exists product_subcategory;
drop table if exists product_category;
drop table if exists product_size;
drop table if exists product_color;
drop table if exists product_source;
drop table if exists products_order;

-- +goose StatementEnd
