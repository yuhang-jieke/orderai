package inits

import (
	"fmt"

	"github.com/yuhang-jieke/orderai/srv/user-server/basic/config"
	"github.com/yuhang-jieke/orderai/srv/user-server/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MysqlInit() {
	conf := GetMysqlConfigFromNacosOrLocal()
	var err error
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf["User"],
		conf["Password"],
		conf["Host"],
		conf["Port"],
		conf["Database"],
	)
	config.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败")
	}
	fmt.Println("数据库连接成功")
	err = config.DB.AutoMigrate(&model.Orders{})
	if err != nil {
		panic("数据表迁移失败")
	}
	fmt.Println("数据表迁移成功")
}
