drop table if exists reserv;
create table reserv(
    id varchar(36) primary key,
    reservation_time timestamp not null,
    reserver_name varchar(255) not null,
    reserver_phone varchar(255) not null,
    table_id varchar(36) not null,
    restaurant_id varchar(36) not null,
    constraint fk_table_id foreign key (table_id) references tables (id)
    );