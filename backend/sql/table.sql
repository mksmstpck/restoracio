drop table if exists tables;

create table if not exists tables (
    id varchar(36) primary key,
    number int not null,
    placement varchar(255) not null,
    max_people int not null,
    is_reserved boolean not null,
    is_occupied boolean not null,
    restaurant_id varchar(36) not null,
    constraint fk_restaurant_id foreign key (restaurant_id) references restaurants (id)
    );