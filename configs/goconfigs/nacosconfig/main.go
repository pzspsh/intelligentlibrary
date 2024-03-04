/*
@File   : main.go
@Author : pan
@Time   : 2023-12-20 17:11:16
*/
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

type NacosConfig struct {
	Ip        string
	Port      uint64
	Namespace string
	DataId    string
	Group     string
}

type Config struct {
	Version  string `json:"version"`
	Username string `json:"username"`
	Password string `json:"password"`
	WebUrl   string `json:"weburl"`
	LogPath  string `json:"logPath"`
}

func (n NacosConfig) FillNacosEnv() NacosConfig {
	nacosIp := os.Getenv("YJWB_NACOS_IP")
	nacosPort := os.Getenv("YJWB_NACOS_PORT")
	fmt.Printf("YJWB_NACOS_IP: %v\n", nacosIp)
	fmt.Printf("YJWB_NACOS_PORT: %v\n", nacosPort)
	if nacosIp != "" {
		n.Ip = nacosIp
	}
	if nacosPort != "" {
		n.Port, _ = strconv.ParseUint(nacosPort, 10, 64)
	}
	return n
}

func getNacosClient(nacosConfig NacosConfig) config_client.IConfigClient {
	// get nacos password
	password := os.Getenv("YJWB_NACOS_PASSWORD")
	if password == "" {
		password = "password"
	}
	//create ServerConfig
	/*
	   sc := []constant.ServerConfig{
	   		*constant.NewServerConfig(nacosConfig.Ip, nacosConfig.Port, constant.WithContextPath("/nacos")),
	   	}
	*/

	sc := []constant.ServerConfig{
		{
			IpAddr:      nacosConfig.Ip,
			ContextPath: "/nacos",
			Port:        nacosConfig.Port,
			Scheme:      "http",
		},
	}

	//create ClientConfig
	/*
		   cc := *constant.NewClientConfig(
				constant.WithNamespaceId(nacosConfig.Namespace),
				constant.WithTimeoutMs(5000),
				constant.WithNotLoadCacheAtStart(true),
				constant.WithLogDir("./tmp/nacos/log"),
				constant.WithCacheDir("./tmp/nacos/cache"),
				constant.WithLogLevel("debug"),
				constant.WithUsername("nacos"),
				constant.WithPassword(password),
			)
	*/
	cc := constant.ClientConfig{
		NamespaceId:         nacosConfig.Namespace,
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "",
		CacheDir:            "",
		LogLevel:            "debug",
		Username:            "nacos",
		Password:            password,
	}

	// create config client
	client, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)

	if err != nil {
		panic(err)
	}

	return client
}

func (x Config) GetConfig(nacosConfig NacosConfig) Config {
	client := getNacosClient(nacosConfig)

	//get config
	content, _ := client.GetConfig(vo.ConfigParam{
		DataId: nacosConfig.DataId,
		Group:  nacosConfig.Group,
	})
	fmt.Println("nacos-config:" + content)

	json.Unmarshal([]byte(content), &x)

	return x
}

func (x Config) LoadConfig() Config {
	nacos := NacosConfig{
		Ip:        "ipconfig",
		Port:      9999,
		Namespace: "test",
		DataId:    "configid",
		Group:     "DEFAULT_GROUP",
	}
	nacos = nacos.FillNacosEnv()
	return x.GetConfig(nacos)
}

func main() {
	var config Config
	config = config.LoadConfig()
	fmt.Printf("%+v\n", config)
}
