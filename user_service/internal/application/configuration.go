package application

import (
	"errors"
	"io/fs"
	"os"
	"reflect"
	"strconv"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

const (
	developmentMode = "development"
	productionMode  = "production"
)

type configuration struct {
	IsProductionMode bool
	EnvMode          string `env:"ENV"`

	ServiceName string `env:"SERVICE_NAME"`

	JWTSecret string `env:"JWT_SECRET"`

	HTTPServerPort string `env:"HTTP_PORT"`
	GRPCServerPort string `env:"GRPC_PORT"`

	ElasticsearchAddress string `env:"ELASTIC_SEARCH"`

	RedisAddress  string `env:"REDIS_HOST"`
	RedisUsername string `env:"REDIS_USERNAME"`
	RedisPassword string `env:"REDIS_PASSWORD"`
	RedisDB       int    `env:"REDIS_DB"`

	MongoConnectionURL string `env:"MONGO_HOST"`

	UsersPostgresConnectionURL string `env:"POSTGRES_HOST"`
}

func newConfiguration() configuration {
	return configuration{}
}

func (conf *configuration) loadConfigurationFromEnvFile() error {
	if err := godotenv.Load("./configs/.env"); err != nil {
		var perr *fs.PathError

		// do nothing if file is absent
		if errors.As(err, &perr) {
			return nil
		}

		return err
	}

	if err := env.Parse(conf); err != nil {
		return err
	}

	conf.IsProductionMode = conf.isProducationMode()

	return nil
}

func (conf configuration) isProducationMode() bool {
	return conf.EnvMode == productionMode
}

func (conf *configuration) loadConfigurationFromEnvironment() error {
	t := reflect.TypeOf(*conf)

	for i := 0; i < t.NumField(); i++ {
		if value, isExists := os.LookupEnv(t.Field(i).Tag.Get("env")); isExists {
			s := reflect.ValueOf(conf).Elem()

			if s.FieldByName(t.Field(i).Name).CanSet() {
				switch t.Field(i).Type.Kind() {
				case reflect.String:
					s.FieldByName(t.Field(i).Name).SetString(value)
				case reflect.Int:
					if in, err := strconv.ParseInt(value, 10, 0); err == nil {
						s.FieldByName(t.Field(i).Name).SetInt(in)
					}
				}
			}
		}
	}

	return nil
}
