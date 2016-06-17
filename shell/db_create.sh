#!/bin/bash

mysql -u root -p <<EOF
create database report;
create user 'reporter'@'localhost' identified by 'reporter';
grant all privileges on * . * to 'reporter'@'localhost';
flush privileges;
EOF

mysql -u reporter -p <<EOF
use report;
create table report_conf(
    id int auto_increment primary key, 
    title varchar(256), 
    description varchar(256),
    table_name varchar(64), 
    handler_key varchar(32), 
    create_time timestamp 
);

create table chartdata_raw(
    id int auto_increment primary key, 
    chart_type varchar(20),
    data_type varchar(20),
    content text,
    create_time timestamp
);

EOF
