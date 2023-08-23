drop table if exists staffs;
create table if not exists staffs (
    id varchar(36) primary key,
    name varchar(255) not null,
    age int not null,
    gender varchar(255) not null,
    phone varchar(255) not null,
    email varchar(255) not null,
    position varchar(255) not null,
    restaurant_id varchar(36) not null,
    constraint fk_restaurant_id foreign key (restaurant_id) references restaurants (id)
    );