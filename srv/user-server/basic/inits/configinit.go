package inits

import (
	"github.com/spf13/viper"
	"github.com/yuhang-jieke/orderai/srv/user-server/basic/config"
)

func ConfigInit() {
	viper.SetConfigFile("C:\\Users\\ZhuanZ\\Desktop\\week3\\orderai\\srv\\config.yaml")
	viper.ReadInConfig()
	viper.Unmarshal(&config.GlobalConf)
}
