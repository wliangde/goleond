
drop table if EXISTS sage_story;
create table sage_story (
  id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '自增id',
  url VARCHAR (200) comment 'url',
  name VARCHAR (200) comment '标题',
  size VARCHAR (200) comment '大小',
  ma VARCHAR (10) ,
  pic VARCHAR (2000) comment 'url',
  tor_url VARCHAR (200) comment 'url',
  download_cnt int,
  UNIQUE index i_url(url),
  index i_download_cnt(download_cnt)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;