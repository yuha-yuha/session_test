use test;
create table users(
    id MEDIUMINT NOT NULL AUTO_INCREMENT,
    user_id varchar(60) not null unique,
    user_name varchar(60) not null,
    user_hash_password varchar(100) not null,
    primary key(id) 
);

