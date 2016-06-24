#!/bin/bash

mysql -u root -p <<EOF
create database report;
create user 'reporter'@'localhost' identified by 'reporter';
grant all privileges on * . * to 'reporter'@'localhost';
flush privileges;
EOF

mysql -u reporter -p <<EOF
use report;
drop table if exists report_conf;
create table if not exists report_conf(
    id int auto_increment primary key, 
    title varchar(256), 
    legend varchar(128),
    description varchar(256),
    table_name varchar(64), 
    data_field varchar(32),
    data_format varchar(10),
    chart_type varchar(10),
    handler_key varchar(32), 
    show_timeline boolean,
    timeline_field varchar(32),
    create_time timestamp 
);

drop table if exists chartdata_raw;
create table chartdata_raw(
    id int auto_increment primary key, 
    chart_type varchar(20),
    data_type varchar(20),
    content text,
    create_time timestamp
);

EOF
