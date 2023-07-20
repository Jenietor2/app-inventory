create database app_inventory;
use app_inventory

create table USERS(
    id int not null auto_increment,
    email varchar(255) not null,
    username varchar(255) not null,
    password varchar(255)not null,
    primary key (id)
)

create table PRODUCTS(
    id int not null auto_increment,
    name varchar(255) not null,
    description varchar(255) not null,
    price int not null,
    created_by int not null,
    primary key (id),
    FOREIGN key (created_by) REFERENCES USERS(id)
)

create table ROLES(
    id int not null auto_increment,
    name varchar(255) not null,
    primary key (id)  
)

create table USER_ROLES(
    id int not null auto_increment,
    user_id int not null,
    role_id int not null,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES USERS(id),
    FOREIGN KEY (role_id) REFERENCES ROLES(id)
)