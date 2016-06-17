#!/bin/bash

mysql -u reporter -preporter <<EOF
use report;
create table if not exists feat10 (id int not null unique auto_increment, raw_data text);
insert into feat10 (raw_data) values("a:0.41 b:0.23 c:0.36");
EOF
