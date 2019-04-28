
-- create database dreba default character set utf8mb4 default collate utf8mb4_general_ci;

SET sql_mode = '';

use dreba;


create table image_infos (
  uuid varchar(50),
  image_data BLOB,
  file_name varchar(200),
  created_at timestamp ,
  updated_at timestamp ,
  primary key (uuid)
);

create table blogs (
  uuid varchar (50),
  title varchar (200),
  tags json,
  context_type varchar(10),
  context MEDIUMTEXT,
  created_at timestamp ,
  updated_at timestamp ,
  primary key (uuid)
);

