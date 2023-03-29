Use blog_service;
create table `blog_tag` (
`id` int(10) unsigned not null auto_increment,
`name` varchar(100) default '' comment '標籤名稱',
/*公用欄位*/
`create_on` int(10) unsigned default '0' comment '建立時間',
`create_by` varchar(100) default '' comment '建立人',
`modified_on` int(10) unsigned default '0' comment '修改時間',
`modified_by` varchar(100) default '' comment '修改人',
`delete_on` int(10) unsigned default '0' comment '刪除時間',
`is_del` tinyint(3) unsigned default '0' comment '是否刪除 0為未刪除 1為已刪除',

`state`  tinyint(3) unsigned default '1' comment '狀態0為禁用,1為啟用',
primary key(`id`)
)engine=InnoDB default charset=utf8mb4 comment='標籤管理';