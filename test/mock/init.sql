create database if not exists testdb;
use testdb;
create table if not exists article (
    title varchar(100) CHARACTER SET utf8mb4,
    body varchar(255) CHARACTER SET utf8mb4,
    userName varchar(50) CHARACTER SET utf8mb4,
    createdAt varchar(200) CHARACTER SET utf8mb4, 
    updatedAt varchar(200) CHARACTER SET utf8mb4, 
    PRIMARY KEY (`title`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;