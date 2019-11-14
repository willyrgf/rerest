package config

import (
	"os"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// KeysConf informations
type KeysConf struct {
	Name     string `mapstructure:"name"`
	Prefix   string `mapstructure:"prefix"`
	Type     string `mapstructure:"type"`
	ReadOnly bool   `mapstructure:"read_only"`
}

// AccessConf informations
type AccessConf map[string]KeysConf

// APIConfig basic config
type APIConfig struct {
	HTTPHost string
	HTTPPort int
}

// Rerest basic config
type Rerest struct {
	API                  *APIConfig
	Access               AccessConf
	RedisHost            string
	RedisPassword        string
	RedisPort            int
	RedisDatabaseData    int
	RedisDatabaseControl int
}

var (
	// RerestConf config variable
	RerestConf        *Rerest
	configFile        string
	defaultConfigFile = "./config.toml"
)

// New config struct
func New() *Rerest {
	api := &APIConfig{}
	access := make(AccessConf)
	return &Rerest{
		API:    api,
		Access: access,
	}
}

func getEnvConfig(env string) (cfg string) {
	cfg = os.Getenv(env)
	return
}

func getDefaultConfig(file string) (fileConfig string) {
	fileConfig = defaultConfigFile
	if file != "" {
		fileConfig = file
	}

	_, err := os.Stat(fileConfig)
	if err != nil {
		fileConfig = ""
	}

	return
}

func viperCfg() {
	configFile = getDefaultConfig(getEnvConfig("REREST_CONF"))
	dir, file := filepath.Split(configFile)
	file = strings.TrimSuffix(file, filepath.Ext(file))
	viper.AddConfigPath(dir)
	viper.SetConfigName(file)
	viper.SetConfigType("toml")
	viper.SetDefault("http.host", "0.0.0.0")
	viper.SetDefault("http.port", 8080)
	viper.SetDefault("redis.host", "127.0.0.1")
	viper.SetDefault("redis.port", 6379)
}

// Parse ReREST configs
func Parse(cfg *Rerest) (err error) {
	err = viper.ReadInConfig()
	if err != nil {
		log.Errorf("config.Parse(): error=%w", err)
		return
	}

	cfg.API.HTTPHost = viper.GetString("http.host")
	cfg.API.HTTPPort = viper.GetInt("http.port")
	cfg.RedisHost = viper.GetString("redis.host")
	cfg.RedisPort = viper.GetInt("redis.port")
	cfg.RedisPassword = viper.GetString("redis.password")
	cfg.RedisDatabaseData = viper.GetInt("redis.database_data")
	cfg.RedisDatabaseControl = viper.GetInt("redis.database_control")

	var ks []KeysConf
	err = viper.UnmarshalKey("access.keys", &ks)
	if err != nil {
		log.Errorf("config.Parse(): error=%w", err)
		return
	}
	access := make(AccessConf)
	for _, k := range ks {
		access[k.Prefix] = k
	}

	cfg.Access = access
	return
}

// Load configuration
func Load() {
	viperCfg()
	RerestConf = New()
	err := Parse(RerestConf)
	if err != nil {
		log.Fatalf("config.Load(): error=%w", err)
	}

	log.Debugf("config.Load(): RerestConf=%+v", RerestConf)
	log.Debugf("config.Load(): RerestConf.Access=%+v", RerestConf.Access)
}
