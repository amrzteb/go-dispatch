package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

var (
	GormPool *gorm.DB
)

func InitMysqlPool()  {
	username := "root"  //账号
	password := "root" //密码
	host := "106.13.177.27" //数据库地址，可以是Ip或者域名
	port := 3307 //数据库端口
	Dbname := "db_ci" //数据库名
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, Dbname)
	var err error
	GormPool, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	//设置数据库连接池参数
	GormPool.DB().SetMaxOpenConns(30)   //设置数据库连接池最大连接数
	GormPool.DB().SetMaxIdleConns(10)   //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭

}