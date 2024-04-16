package portal_sync

type Config struct {
	Database *DatabaseConfig
	Portal   *PortalConfig
}

type DatabaseConfig struct {
	Server   string
	Database string
	User     string
	Password string
}

type PortalConfig struct {
	Server   string
	User     string
	Password string
	Timeout  int
}
