package config

import (
	"os"
	"time"

	_ "github.com/lib/pq" // postgres driver don`t delete
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type ServerConfig struct {
	Name     string `mapstructure:"name"`
	Port     string `mapstructure:"port"`
	BasePath string `mapstructure:"base_path"`
	TestMode bool   `mapstructure:"test_mode"`
	Log      struct {
		Level  string `mapstructure:"level"`
		Format string `mapstructure:"format"`
	} `mapstructure:"log"`
}

type JWTConfig struct {
	AccessToken struct {
		SecretKey     string        `mapstructure:"secret_key"`
		TokenLifetime time.Duration `mapstructure:"token_lifetime"`
	} `mapstructure:"access_token"`
	ServiceToken struct {
		SecretKey string `mapstructure:"secret_key"`
	} `mapstructure:"service_token"`
}

type KafkaConfig struct {
	Brokers      []string      `mapstructure:"brokers"`
	Topic        string        `mapstructure:"topic"`
	GroupID      string        `mapstructure:"group_id"`
	DialTimeout  time.Duration `mapstructure:"dial_timeout"`
	ReadTimeout  time.Duration `mapstructure:"read_timeout"`
	WriteTimeout time.Duration `mapstructure:"write_timeout"`
	RequiredAcks string        `mapstructure:"required_acks"`
}

type AwsConfig struct {
	Region      string `mapstructure:"region"`
	BucketName  string `mapstructure:"bucket_name"`
	AccessKeyID string `mapstructure:"access_key_id"`
	AccessKey   string `mapstructure:"access_key"`
}

type SwaggerConfig struct {
	Enabled bool   `mapstructure:"enabled"`
	URL     string `mapstructure:"url"`
	Port    string `mapstructure:"port"`
}

type DatabaseConfig struct {
	SQL struct {
		URL string `mapstructure:"url"`
	} `mapstructure:"sql"`

	Redis struct {
		Addr     string `mapstructure:"addr"`
		Password string `mapstructure:"password"`
		DB       int    `mapstructure:"db"`
		Lifetime int    `mapstructure:"lifetime"`
	} `mapstructure:"redis"`
}

type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	JWT      JWTConfig      `mapstructure:"jwt"`
	Kafka    KafkaConfig    `mapstructure:"kafka"`
	Aws      AwsConfig      `mapstructure:"aws"`
	Database DatabaseConfig `mapstructure:"database"`
	Swagger  SwaggerConfig  `mapstructure:"swagger"`
	Log      *logrus.Logger
}

func LoadConfig() (Config, error) {
	configPath := os.Getenv("KV_VIPER_FILE")
	if configPath == "" {
		return Config{}, errors.New("KV_VIPER_FILE env var is not set")
	}
	viper.SetConfigFile(configPath)

	if err := viper.ReadInConfig(); err != nil {
		return Config{}, errors.Errorf("error reading config file: %s", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return Config{}, errors.Errorf("error unmarshalling config: %s", err)
	}

	return config, nil
}
