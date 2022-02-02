CREATE DATABASE IF NOT EXISTS docker CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
CREATE USER IF NOT EXISTS 'docker'@'%' IDENTIFIED BY 'docker';
GRANT ALL PRIVILEGES ON docker.* TO 'docker'@'%';

create table my_testdb.users (id int, name varchar(100));
insert into my_testdb.users values (1, 'ichiro');
insert into my_testdb.users values (2, 'jiro');
insert into my_testdb.users values (3, 'saburo');
insert into my_testdb.users values (4, 'shiro');
insert into my_testdb.users values (5, 'goro');
insert into my_testdb.users values (6, 'rokuro');
insert into my_testdb.users values (7, 'shichiro');
insert into my_testdb.users values (8, 'hachiro');
insert into my_testdb.users values (9, 'kuro');
insert into my_testdb.users values (10, 'juro');

FLUSH PRIVILEGES;