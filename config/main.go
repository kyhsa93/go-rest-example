package config

// Config config stcut
type Config struct {
	Swagger  *Swagger
	Auth     *Auth
	Server   ServerConfigInterface
	Database *Database
	Redis    *Redis
	Email    *Email
	AWS      *AWS
}

// InitializeConfig initialize config
func InitializeConfig() *Config {
	return &Config{
		Server:   NewServerConfig(),
		Database: NewDatabase(),
		Swagger:  NewSwagger(),
		Auth:     NewAuth(),
		Redis:    NewRedis(),
		Email:    NewEmail(),
		AWS:      NewAWS(),
	}
}
