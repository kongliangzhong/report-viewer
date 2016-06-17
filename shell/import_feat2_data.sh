#!/bin/bash

mysql -u reporter -preporter <<EOF
use report;
create table if not exists feat2 (id int not null unique, raw_miu int, smooth_miu int, distribution text);
load data infile '/usr/local/klzhong/workspace/creditease/gitrepos/report-viewer/raw_data/a.txt' into table feat2 
  fields terminated by '\t'
  lines terminated by '\n' 
  ignore 1 lines; ## ignore 1 line at start of data file.
EOF
