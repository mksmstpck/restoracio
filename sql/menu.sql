drop table if exists menus;
create table if not exists menus (
    id varchar(36) primary key,
    name varchar(255) not null,
    description varchar(1024),
    qrcode bytea,
    restaurant_id varchar(36)
    );