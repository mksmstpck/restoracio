drop table if exists restaurant;

create table if not exists restaurant (
    id varchar(36) primary key,
    name varchar(255) not null,
    location varchar(255) not null,
    staff_ids varchar(36)[] not null,
    dish_ids varchar(36)[] not null,
    table_ids varchar(36)[] not null,
    admin_id varchar(36)[] not null
    );