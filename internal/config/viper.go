package config

import (
	"errors"

	"github.com/spf13/viper"
)

const (
	EnvDevelop = "develop"
	EnvTest    = "test"
	EnvRelease = "release"
)

type Config struct {
	SrvEnv  string
	SrvName string
	SrvAddr string
	*viper.Viper
}

func New(env, name, addr, configPath string) (*Config, error) {
	if env != EnvDevelop && env != EnvTest && env != EnvRelease {
		return nil, errors.New("unknown env variable")
	}
	cfg := &Config{
		env,
		name,
		addr,
		viper.New(),
	}

	cfg.SetConfigFile(configPath)
	if err := cfg.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := cfg.MergeConfigMap(cfg.GetStringMap("common")); err != nil {
		return nil, err
	}
	if err := cfg.MergeConfigMap(cfg.GetStringMap(cfg.SrvEnv)); err != nil {
		return nil, err
	}

	if n := cfg.GetString("srvName"); len(n) > 0 {
		cfg.SrvName = n
	}
	if a := cfg.GetString("srvAddr"); len(a) > 0 {
		cfg.SrvAddr = a
	}

	return cfg, nil
}
