package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Cfg struct {
	Db   Dbcfg
	Auth AuthCfg
}

type Dbcfg struct {
	Dsn string
}

type AuthCfg struct {
	Secret string
}

func LoadConfig() *Cfg {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env, using default cfg")
	}
	return &Cfg{
		Db: Dbcfg{
			Dsn: os.Getenv("DSN"),
		},
		Auth: AuthCfg{
			Secret: os.Getenv("TOKEN"),
		},
	}
}
