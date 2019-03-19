#!/bin/bash
HOST="127.0.0.1"
USER="root"
PWD="10086"
DB_NAME="lianjia"

#获取当前时间 0227
TIME=$(date "+%m%d")  

TABLE="house"
NEW_TABLE=$TABLE"_"$TIME

echo $NEW_TABLE
#备份数据库
mysql -u$USER -p$PWD -h$HOST $DB_NAME -e "CREATE TABLE $NEW_TABLE LIKE $TABLE; INSERT INTO $NEW_TABLE SELECT * FROM $TABLE; TRUNCATE $TABLE";

#执行爬虫脚本
base_dir="/root/go/src/github.com/wliangde/goleond/example/spider_lianjia"
nohup $base_dir/spider_lianjia -xq=false > $base_dir/log/house_$TIME.log &

