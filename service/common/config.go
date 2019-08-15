package common

import (
	"github.com/gomodule/redigo/redis"
	"log"
)

const CONFIG_FILE = "config.ini"

func ConnectRedis() redis.Conn {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Println(err)
	}
	return conn
}

//var m *AppConfig
//var once sync.Once // go 的并发包 sync 里的工具类，保证某个操作只能执行一次
//type AppConfig struct {
//	MySql struct{
//		User string `yaml:"user"` // yaml 的元数据，定义了怎么解析 yaml 文件
//		Host string `yaml:"hostname"`
//		Port string `yaml:"port"`
//		Password string `yaml:"password"`
//		DBName string `yaml:"database"`
//	}
//}
//
//// 单例模式
//func CfgInstance() *AppConfig {
//	once.Do(func() {
//		m,_ = loadConf()
//	})
//	return m
//}
//
//func loadConf() (*AppConfig, error) {
//	localConfPath := "../config.yml"
//
//	conf := &AppConfig{}
//	yamlFile, err := ioutil.ReadFile(localConfPath)
//	if err != nil {
//		println("Error! Yaml file IOException: %s", err.Error())
//		os.Exit(1)
//	}
//	// 从字节数组里反序列化成一个 conf 实例
//	if err := yaml.Unmarshal(yamlFile, conf); err != nil {
//		println("Error! Yaml unmarshal exception: %s", err.Error())
//		os.Exit(1)
//	}
//	return conf, nil
//}
