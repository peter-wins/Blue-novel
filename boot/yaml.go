package boot

import (
	"fmt"
	"github.com/spf13/viper"
)

// yaml配置初始化
func yamlInit(){
	// 加载yaml翻译配置文件
	viper.SetConfigName("messages.zh_CN.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./translations")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error global file: %s \n", err))
	}
	fmt.Println("yaml翻译配置初始化完成...")
}