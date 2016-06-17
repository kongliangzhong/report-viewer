create database if not exists stat;
use stat;

drop table if exists feat_distribution_stat_info;
create table if not exists feat_distribution_stat_info (
	`id` bigint(20) NOT NULL AUTO_INCREMENT comment 'pk',
	`feat_id` MEDIUMINT NOT NULL comment 'feat id',
	`distribution_kv` MEDIUMTEXT NOT NULL comment 'k1:v1 k2:v2',
	`outlier_candidates` TEXT NULL DEFAULT NULL comment 'o1,o2',
	`high_density_candidates` TEXT NULL DEFAULT NULL comment 'd1,d2',
	`mean` float DEFAULT NULL comment 'mean value of feat of continuous',
	`stddev` float DEFAULT NULL comment 'standard deviation  of feat of continuous',
	`dt` int NOT NULL DEFAULT 19700101 comment 'end day of stat time interval',
	PRIMARY KEY (`id`),
	UNIQUE KEY `dtfeat` (`dt`, `feat_id`)
) ENGINE=InnoDB DEFAULT CHARACTER SET=utf8;

drop table if exists feat_map;
create table if not exists feat_map (
	id int NOT NULL AUTO_INCREMENT comment 'feat id',
	feat_name VARCHAR(32) NOT NULL comment 'feat name',
	dimension VARCHAR(20) NULL DEFAULT NULL comment 'dimensional unit',
	feat_type ENUM('discrete', 'continuous') NOT NULL comment 'feat type',
	extra_info VARCHAR(3000) DEFAULT NULL comment 'extra info',
	PRIMARY KEY (`id`),
	UNIQUE KEY `featname` (`feat_name`)
) ENGINE=InnoDB DEFAULT CHARACTER SET=utf8;

insert into feat_map(feat_name, dimension, feat_type, extra_info) values ('month_repay_amount_min', '万元', 'continuous', '用户月最低还款额');
insert into feat_distribution_stat_info(feat_id,
                                        distribution_kv,
					outlier_candidates,
					high_density_candidates,
					mean,
					stddev,
					dt) values 
					(
					1,
					'0:100000,1:100,2:200,3:100,10000:1',
					'0',
					'10000',
					0.54,
					0.1,
					20160601);
