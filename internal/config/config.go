package config

import (
	"changeme/internal/consts"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"time"
)

type Configs struct {
	App app    `yaml:"app"`
	Log log    `yaml:"log"`
	Db  sqlite `yaml:"sqlite"`
}

type sqlite struct {
	FilePath        string `yaml:"filePath"`
	MaxIdleCons     int    `yaml:"maxIdleCons"`
	MaxOpenCons     int    `yaml:"maxOpenCons"`
	ConnMaxLifetime int    `yaml:"connMaxLifetime"`
}

type app struct {
	Level     string        `yaml:"level"`
	GWTimeOut time.Duration `yaml:"GWTimeOut"`
}

type log struct {
	FilePath string `yaml:"filePath"`
}

var Conf Configs

func Init() (err error) { //初始化config全局变量
	var bs []byte
	if bs, err = ioutil.ReadFile(consts.ConfigPath); err != nil {
		return fmt.Errorf("读取config文件失败: %s", err)
	}

	if err = yaml.UnmarshalStrict(bs, &Conf); err != nil {
		return fmt.Errorf("yaml序列化失败: %s", err)
	}
	return nil
}
