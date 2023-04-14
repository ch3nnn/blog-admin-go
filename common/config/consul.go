package config

import (
	"errors"
	"github.com/ch3nnn/blog-admin-go/common/encoding"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"

	"github.com/hashicorp/consul/api"
)

type ConsulConfig struct {
	Consul consul.Conf
}

// Validate validates c.
func (c ConsulConfig) Validate() error {
	if len(c.Consul.Host) == 0 {
		return errors.New("empty consul hosts")
	}
	if len(c.Consul.Key) == 0 {
		return errors.New("empty consul key")
	}
	if c.Consul.TTL == 0 {
		c.Consul.TTL = 20
	}

	return nil
}

// NewClient create new client
func (c ConsulConfig) NewClient() *api.Client {
	client, err := api.NewClient(&api.Config{Scheme: "http", Address: c.Consul.Host, Token: c.Consul.Token})
	if err != nil {
		logx.Must(err)
		return nil
	}
	return client
}

// LoadYAMLConf load config from consul kv
func LoadYAMLConf(client *api.Client, key string, v interface{}) {
	kv := client.KV()

	data, _, err := kv.Get(key, nil)
	b, err := encoding.YamlToJson(data.Value)
	if err != nil {
		logx.Must(err)
	}
	err = conf.LoadFromJsonBytes(b, v)
	logx.Must(err)
}

// LoadJSONConf load config from consul kv
func LoadJSONConf(client *api.Client, key string, v interface{}) {
	kv := client.KV()

	data, _, err := kv.Get(key, nil)
	err = conf.LoadFromJsonBytes(data.Value, v)
	logx.Must(err)
}
