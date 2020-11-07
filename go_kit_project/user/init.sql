create schema if not exists test_db;
create table if not exists test_db.user
(
    id         bigint auto_increment
        primary key,
    username   varchar(100)                        not null,
    password   varchar(100)                        not null,
    email      varchar(100)                        not null,
    created_at timestamp default CURRENT_TIMESTAMP not null
);