package config

var GlobCfg = Config{}

type Config struct {
	Prompt string `toml:"prompt"`
}
