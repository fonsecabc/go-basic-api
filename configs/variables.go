package configs

import (
	"github.com/go-chi/jwtauth"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Variables struct {
	DBDriver     string `envconfig:"DB_DRIVER" required:"true"`
	DBHost       string `envconfig:"DB_HOST" required:"true"`
	DBPort       string `envconfig:"DB_PORT" required:"true"`
	DBUser       string `envconfig:"DB_USER" required:"true"`
	DBPass       string `envconfig:"DB_PASS" required:"true"`
	DBName       string `envconfig:"DB_NAME" required:"true"`
	JWTSecret    string `envconfig:"JWT_SECRET" required:"true"`
	JWTExpiresIn int    `envconfig:"JWT_EXPIRES_IN" required:"true"`
	TokenAuth    *jwtauth.JWTAuth
}

var Env Variables

func LoadVariables() error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}

	err = envconfig.Process("", &Env)
	if err != nil {
		return err
	}

	Env.TokenAuth = jwtauth.New("HS256", []byte(Env.JWTSecret), nil)

	return nil
}
