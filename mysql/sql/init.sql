CREATE DATABASE IF NOT EXISTS docker CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
CREATE USER IF NOT EXISTS 'docker'@'%' IDENTIFIED BY 'docker';
GRANT ALL PRIVILEGES ON docker.* TO 'docker'@'%';

create table my_testdb.users (id int not null primary key auto_increment, name varchar(100) not null, password varchar(10) not null);
insert into my_testdb.users values (1, 'ichiro', 'password');
insert into my_testdb.users values (2, 'jiro', 'password');
insert into my_testdb.users values (3, 'saburo', 'password');
insert into my_testdb.users values (4, 'shiro', 'password');
insert into my_testdb.users values (5, 'goro', 'password');
insert into my_testdb.users values (6, 'rokuro', 'password');
insert into my_testdb.users values (7, 'shichiro', 'password');
insert into my_testdb.users values (8, 'hachiro', 'password');
insert into my_testdb.users values (9, 'kuro', 'password');
insert into my_testdb.users values (10, 'juro', 'password');

FLUSH PRIVILEGES;