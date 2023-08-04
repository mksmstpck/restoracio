drop table if exists restaurants;

create table if not exists restaurants (
    id varchar(36) primary key,
    name varchar(255) not null,
    location varchar(255) not null,
    admin_id varchar(36) not null,
    constraint fk_admin_id foreign key (admin_id) references admins (id)
    );