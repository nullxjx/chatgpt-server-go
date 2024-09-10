package config

import (
	"encoding/json"
	"os"
)

// Config 服务配置
type Config struct {
	HttpPort  int    `json:"http_port"`
	OpenAIKey string `json:"openai_key"`
	BaseURL   string `json:"base_url"`
}

// storeDefault 设置默认值。
func (c *Config) storeDefault() {
	c.HttpPort = 8080
	c.BaseURL = "https://api.openai.com/v1"
}

// parseFromFile 从文件解析配置。
func (c *Config) parseFromFile(path string) error {
	ds, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(ds, c); err != nil {
		return err
	}
	return nil
}

// New 返回一个新的配置对象。
func New(path string) (*Config, error) {
	var c Config
	c.storeDefault()
	if err := c.parseFromFile(path); err != nil {
		return nil, err
	}
	return &c, nil
}
