package env

import "os"

type EnvKey string

func (k EnvKey) GetValue() string {
	return os.Getenv(string(k))
}

const (
	DbDsn     EnvKey = "DB_DSN"
	DriverDb  EnvKey = "DRIVER_DB"
	SecretJWT EnvKey = "SECRETJWT"
)
