#!/bin/bash

mysql -u reporter -preporter <<EOF
use report;
##--insert into report_conf(title, description, table_name, handler_key) values('feat2 distribution histogram', '' , 'feat2', 'feat2-bar');
insert into report_conf(title, description, table_name, handler_key) values('feat10 distribution piechart', '' , 'feat10', 'feat10-pie');
EOF
