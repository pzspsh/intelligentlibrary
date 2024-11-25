/*
@File   : main.go
@Author : pan
@Time   : 2024-11-20 16:57:58
*/
package main

import (
	"encoding/json"
	"fmt"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

type TargetConfig struct { // 例如：你要获取的目标配置
	Mysql    Mysql    `json:"mysql,omitempty"`
	Postgres Postgres `json:"postgres,omitempty"`
}

type Mysql struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}

type Postgres struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}

type NacosConfig struct {
	IP        string `json:"ip"`
	Port      uint64 `json:"port"`
	Namespace string `json:"namespace"`
	Password  string `json:"password"`
	DataId    string `json:"dataid"`
	Group     string `json:"group"`
}

func (n *NacosConfig) LoadConfig() (*TargetConfig, error) {
	var err error
	var config *TargetConfig
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(n.IP, n.Port, constant.WithContextPath("/nacos")),
	}
	/* cc := *constant.NewClientConfig(
		constant.WithNamespaceId(n.Namespace),
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir("./tmp/nacos/log"),
		constant.WithCacheDir("./tmp/nacos/cache"),
		constant.WithLogLevel("debug"),
		constant.WithUsername("nacos"),
		constant.WithPassword(n.Password),
	) */
	cc := constant.ClientConfig{
		NamespaceId:         n.Namespace,
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "", // 不写日志
		CacheDir:            "", // 不写缓存
		LogLevel:            "debug",
		Username:            "nacos",
		Password:            n.Password,
		// LogDir:              "./tmp/nacos/log",
		// CacheDir:            "./tmp/nacos/cache",
	}
	client, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		return config, err
	}
	content, err := client.GetConfig(vo.ConfigParam{
		DataId: n.DataId,
		Group:  n.Group,
	})
	if err != nil {
		return config, err
	}
	if err := json.Unmarshal([]byte(content), &config); err != nil {
		return config, err
	}
	return config, err
}

func main() {
	nacos := &NacosConfig{
		IP:        "your-ip",
		Port:      8848, // your-port
		Namespace: "your-namespace",
		Password:  "your-password",
		DataId:    "your-dataid",
		Group:     "your-group",
	}
	config, err := nacos.LoadConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(config)
}
