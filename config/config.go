/*
 * @Author: lisheng
 * @Date: 2022-10-11 14:56:11
 * @LastEditTime: 2023-02-07 13:35:07
 * @LastEditors: lisheng
 * @Description: 读取配置
 * @FilePath: /jf-go-kit/config/config.go
 */
package config

import (
	"fmt"

	"github.com/hatlonely/go-kit/bind"
	"github.com/hatlonely/go-kit/config"
	"github.com/hatlonely/go-kit/flag"
	"github.com/hatlonely/go-kit/refx"
	"gopkg.in/yaml.v2"
)

var C Config

type mysql struct {
	Addr string `mapstructure:"addr"`
	Port int    `mapstructure:"port"`
	User string `mapstructure:"user"`
	Pwd  string `mapstructure:"pwd"`
	DB   string `mapstructure:"db"`
}

type clickhouse struct {
	Addr string `mapstructure:"addr"`
	Port int    `mapstructure:"port"`
	User string `mapstructure:"user"`
	Pwd  string `mapstructure:"pwd"`
	DB   string `mapstructure:"db"`
}

type pg struct {
	Addr string `mapstructure:"addr"`
	Port int    `mapstructure:"port"`
	User string `mapstructure:"user"`
	Pwd  string `mapstructure:"pwd"`
	DB   string `mapstructure:"db"`
}

type mongo struct {
	Addr string `mapstructure:"addr"`
	Port int    `mapstructure:"port"`
	User string `mapstructure:"user"`
	Pwd  string `mapstructure:"pwd"`
	DB   string `mapstructure:"db"`
}

type redis struct {
	Addr string `mapstructure:"addr"`
	Pwd  string `mapstructure:"pwd"`
}

type tdengine struct {
	Addr string `mapstructure:"addr"`
	Port int    `mapstructure:"port"`
	User string `mapstructure:"user"`
	Pwd  string `mapstructure:"pwd"`
	DB   string `mapstructure:"db"`
}

type Config struct {
	flag.Options
	Usingtd string `json:"usingtd"`
	Log     struct {
		Level string `mapstructure:"level"`
		Path  string `mapstructure:"path"`
		Name  string `mapstructure:"name"`
	} `json:"log"`
	Server struct {
		Listen string `mapstructure:"listen"`
	} `json:"server"`
	Mysql      map[string]mysql      `json:"mysql"`
	Clickhouse map[string]clickhouse `json:"clickhouse"`
	Pg         map[string]pg         `json:"pg"`
	Mongo      map[string]mongo      `json:"mongo"`
	Tdengine   map[string]tdengine   `json:"tdengine"`
	Redis      map[string]redis      `json:"redis"`
	Rabbitmq   map[string]string     `json:"rabbitmq"`
}

func (c Config) String() string {
	if b, err := yaml.Marshal(c); err == nil {
		return string(b)
	}
	return ""
}

func (c Config) show() {
	config := fmt.Sprintf(`Config
--------------------------------------------------------------
%v--------------------------------------------------------------
`, c)
	fmt.Println(config)
}

func Initialize(configPath string) {
	refx.Must(flag.Struct(&C, refx.WithCamelName()))
	refx.Must(flag.Parse(flag.WithJsonVal()))
	if C.Help {
		fmt.Println(flag.Usage())
		return
	}
	if configPath == "" {
		C.ConfigPath = "config/app.json"
	} else {
		C.ConfigPath = configPath
	}
	cfg, err := config.NewConfigWithSimpleFile(C.ConfigPath)
	refx.Must(err)
	refx.Must(bind.Bind(&C, []bind.Getter{flag.Instance(),
		bind.NewEnvGetter(bind.WithEnvPrefix("GO")), cfg}, refx.WithCamelName()))
	C.show()
}
