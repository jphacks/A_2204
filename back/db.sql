use dietapp;
create table user(
    ID INT AUTO_INCREMENT,
    auth0_id varchar,
    height float,
    birthday data,
    PRIMARY KEY (ID)
);
create table user_meals(
    ID INT AUTO_INCREMENT,
    user_id int,
    name varchar,
    cal int,
    at data,
    PRIMARY KEY (ID)
);
create table user_weight(
    ID INT AUTO_INCREMENT,
    user_id int,
    weight float,
    at data,
    PRIMARY KEY (ID),
);
create table characters(
    user_id int,
    name varchar,
    level int,
    weight float,
    at data,
    exp int,
    PRIMARY KEY (ID),
);
