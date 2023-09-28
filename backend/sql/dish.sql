drop table if exists dishes;
create table if not exists dishes (
    id varchar(36) primary key,
    name varchar(255) not null,
    type varchar(255) not null,
    category varchar(255) not null,
    price int not null,
    curency varchar(255) not null,
    mass_grams int not null,
    ingredients varchar(255)[],
    description varchar(255),
    menu_id varchar(36),
    constraint fk_menu_id foreign key (menu_id) references menus (id)
    );