use dietapp;
create table users(
    ID INT AUTO_INCREMENT,
    auth0_id varchar,
    height float,
    birthday data,
    PRIMARY KEY (id)
);
create table user_meals(
    ID INT AUTO_INCREMENT,
    user_id int,
    name varchar,
    cal int,
    at data,
    PRIMARY KEY (id)
);
create table user_weights(
    ID INT AUTO_INCREMENT,
    user_id int,
    weight float,
    at data,
    PRIMARY KEY (id),
);
create table characters(
    user_id int,
    name varchar,
    level int,
    weight float,
    at data,
    exp int,
    PRIMARY KEY (id),
);
