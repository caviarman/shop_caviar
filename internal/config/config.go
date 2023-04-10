package config

type Config struct {
	DataBase
	HTTP
}

type DataBase struct {
	URL                string `env-required:"true" env:"DATABASE_URL"`
	MaxPoolSize        int32  `env-default:"1"     env:"DATABASE_MAX_POOL_SIZE"`
	MigrateWithIndexes bool   `env:"MIGRATE_WITH_INDEXES"`
}

type HTTP struct {
	Port        string `env-default:"443"  env:"PORT"`
	PermitLimit int    `env-default:"30"   env:"HTTP_PERMIT_LIMIT"`
}
