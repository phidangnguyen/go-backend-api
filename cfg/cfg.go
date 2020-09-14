package cfg

type ServiceCfg struct {
	// Server generate config
	ServerMod string
	Port      string

	// Jwt token configuration
	JWTSecretKey string

	// DB
	PostgresURI string
}
