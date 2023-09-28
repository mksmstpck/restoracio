drop table if exists admins;

create table if not exists admins (
    id varchar(36) primary key,
    name varchar(255) not null,
    email varchar(255) not null,
    password varchar(255) not null,
    restaurant_id varchar(36)
    );