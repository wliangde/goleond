#!/bin/sh
NO_CONF=0
if [ "$1" = "nc" ]; then
	NO_CONF=1
fi

OLD_CONFIG_DATA=../../data  #配置表目录

#更新代码
echo "更新代码......"
svn up
#重新编译
echo "编译......"
make clean
make

mkdir -p pack 

cp gmweb stop.sh  pack/

#使用发布的配置
cp start.pub.sh pack/start.sh

#创建pack目录
mkdir -p pack/log

#更新配置表
echo "更新策划表......"
svn up $OLD_CONFIG_DATA

#拷贝配置表
#cp $OLD_CONFIG_DATA $NEW_CONFIG_DATA -r
echo "拷贝策划配置表......"
rsync --quiet -avz --delete --exclude='.svn/'  $OLD_CONFIG_DATA/ pack/config_data
#cp conf pack/ -r

echo "拷贝gmweb数据......"
#拷贝views
rsync --quiet -avz --delete --exclude='.svn/' views pack/
#拷贝static
rsync --quiet -avz --delete --exclude='.svn/' static pack/
#拷贝data
rsync --quiet -avz --delete --exclude='.svn/' data pack/

#拷贝conf
if [ "$1" = "nc" ]
then
	echo "不拷贝配置"
else
	echo "拷贝gmweb配置..."
	rsync --quiet -avz --delete --exclude='.svn/' conf pack/
	
	#覆盖发布配置
	mv pack/conf/app.pub.conf pack/conf/app.conf
fi
