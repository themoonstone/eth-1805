-- create database mydb charset=utf8;
-- use mydb;

create table person(
  id int primary key auto_increment,
  name varchar(20),
  age int,
  rmb float,
  gender bool,
  birthday date
);

insert into person(name,age,rmb,gender,birthday) values
("张全蛋",25,123456,true,19900921),
("张半丹",13,0.5,true,20120921),
("张没蛋",1,0,true,20170921);