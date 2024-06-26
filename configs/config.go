package configs

import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

type Conf struct {
	// App
	WebServerPort string `mapstructure:"WEB_SERVER_PORT"`

	// Database
	DBDriver   string `mapstructure:"DB_DRIVER"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBName     string `mapstructure:"DB_NAME"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`

	// JWT
	JWTSecret    string `mapstructure:"JWT_SECRET"`
	JWTExpiresIn int    `mapstructure:"JWT_EXPIRES_IN"`
	TokenAuth    *jwtauth.JWTAuth
}

func LoadConfig(path string) (*Conf, error) {
	var cfg *Conf

	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	cfg.TokenAuth = jwtauth.New("HS256", []byte(cfg.JWTSecret), nil)

	return cfg, err
}
