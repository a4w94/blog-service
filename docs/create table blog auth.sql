Use blog_service;
create table `blog_auth` (
`id` int(10) unsigned not null auto_increment,
`app_key` varchar(20) default '' comment 'Key',
`app_secret` varchar(50) default '' comment 'Secret',
/*公用欄位*/
`created_on` int(10) unsigned default '0' comment '建立時間',
`created_by` varchar(100) default '' comment '建立人',
`modified_on` int(10) unsigned default '0' comment '修改時間',
`modified_by` varchar(100) default '' comment '修改人',
`deleted_on` int(10) unsigned default '0' comment '刪除時間',
`is_del` tinyint(3) unsigned default '0' comment '是否刪除 0為未刪除 1為已刪除',
primary key (`id`) using btree
)engine=InnoDB default charset=utf8mb4 comment='認證管理';