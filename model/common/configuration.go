package common

type Configuration struct {
	Token    TokenConfiguration
	Logger   LoggerConfiguration
	Database DatabaseConfiguration
	Cache    CacheConfiguration
	Server   ServerConfiguration
}

type TokenConfiguration struct {
	Auth string
}

type LoggerConfiguration struct {
}

type DatabaseConfiguration struct {
	DBUser            string
	DBPassword        string
	DBHost            string
	DBPort            string
	DBName            string
	DBMaxIdleConns    int
	DBMaxOpenConns    int
	DBConnMaxLifetime int
	DBConnMaxIdleTime int
}

type CacheConfiguration struct {
	CacheHost     string
	CachePort     string
	CachePassword string
}

type ServerConfiguration struct {
	ServerHost string
	ServerPort string
}
