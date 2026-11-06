package config

import (
	"errors"
	"log"
	"os"

	"go.yaml.in/yaml/v3"
)

type Config struct {
	Redis RedisConfig `yaml:"Redis"`
}
type RedisConfig struct {
	Hostname string `yaml:"Hostname"`
	Port     string `yaml:"Port"`
	Password string `yaml:"Password"`
	Db       int    `yaml:"Db"`
}

func InitConfig() (*Config, error) {
	cfg := DefaultConfig()
	byte, err := os.ReadFile("config.yaml")
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		log.Printf("读取配置文件失败:%s", err)
		return nil, err
	}
	if errors.Is(err, os.ErrNotExist) {
		log.Println("未找到配置文件,使用默认配置!")
		return cfg, nil
	}
	if err = yaml.Unmarshal(byte, &cfg); err != nil {
		log.Printf("反序列化配置文件时发生错误:%s", err)
		return nil, err
	}
	return cfg, nil
}

func DefaultConfig() *Config {
	return &Config{
		Redis: RedisConfig{
			Hostname: "localhost",
			Port:     "6379",
			Password: "",
			Db:       0,
		},
	}
}
