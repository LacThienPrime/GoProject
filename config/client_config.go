package config

type ClientConfig struct {
	Mode string `json:"mode"`
	Name string `json:"name"`
}

func (c *Config) clientUser() *ClientConfig {
	a := c.clientAssets()

	cfg := &ClientConfig{
		Mode: a.AppCss,
	}

	return cfg
}

func (c *Config) ClientSession() (cfg *ClientConfig) {
	cfg = c.clientUser()
	return cfg
}
