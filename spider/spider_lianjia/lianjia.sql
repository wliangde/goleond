

drop table if EXISTS xiao_qu;
create table xiao_qu (
  id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '自增id',
  xiao_qu_id bigint comment '小区链家id',
  name VARCHAR(200) COMMENT '小区名',
  price int comment '均价',
  sell_cnt int comment '出售数量',
  sold_cnt int comment '90天出售数量',
  url VARCHAR (200) comment '小区url',
  area VARCHAR (40) comment '所属区域',
  UNIQUE index i_area_id(xiao_qu_id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

drop table if EXISTS house;
create table house (
  id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '自增id',
  house_id bigint comment '房屋链家id',
  url VARCHAR (200) comment '房屋url',
  title VARCHAR (200) comment '标题',
  total_price int comment '总价',
  price int comment '单价',
  xiao_qu_id bigint comment '',
  xiao_qu_name VARCHAR (200),
  xiao_qu_url VARCHAR (200),
  area VARCHAR (200),
  mian_ji int comment '面积',
  hu_xing VARCHAR (100) comment '户型',
  zhuang_xiu VARCHAR (100) comment '装修',
  flood VARCHAR (200) comment '楼层',
  build_time VARCHAR (200) comment '建造时间',
  house_info VARCHAR (300) comment '房屋信息',
  follow int  comment '关注人数',
  look int comment '带看人数',
  release_time VARCHAR (100) comment '发布时间',
  tag VARCHAR (100),

  UNIQUE index i_house_id(house_id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;