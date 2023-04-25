Use blog_service;
create table `blog_article`(
`id` int(10) unsigned not null auto_increment,
`title` varchar(100) default '' comment '文章標題',
`desc` varchar(255) default '' comment '文章簡述',
`cover_image_url` varchar(255) default '' comment '封面圖片位址',
`content` longtext comment '封面圖片位址',

/*公用欄位*/
`created_on` int(10) unsigned default '0' comment '建立時間',
`created_by` varchar(100) default '' comment '建立人',
`modified_on` int(10) unsigned default '0' comment '修改時間',
`modified_by` varchar(100) default '' comment '修改人',
`delete_on` int(10) unsigned default '0' comment '刪除時間',
`is_del` tinyint(3) unsigned default '0' comment '是否刪除 0為未刪除 1為已刪除',

`state`  tinyint(3) unsigned default '1' comment '狀態0為禁用,1為啟用',
primary key(`id`)
)engine=InnoDB default charset=utf8mb4 comment='文章管理';