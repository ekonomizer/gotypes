package gotypes

// Log - параметры логера
type Log struct {
	Output   string `env:"LOG_OUTPUT"`
	MinLevel string `env:"LOG_LEVEL"`
	Syslog   struct {
		Tag string `env:"SYSLOG_TAG"`
		Network string `env:"RSYSLOG_NETWORK"`
		Address string `env:"RSYSLOG_ADDRESS"`
	}
}

// Rabbit - параметры подключения к RabbitMQ
type Rabbit struct {
	URL               string `env:"RABBIT_URL"`
	Channel           string `env:"RABBIT_CHANNEL"`
	AutoAck           bool   `env:"RABBIT_AUTO_ACK"`
	ReconnectAttempts int    `env:"RABBIT_RECONNECT_ATTEMPTS"`
}

// HTTPServer - параметры HTTP-сервера
type HTTPServer struct {
	Host string `env:"HTTP_HOST"`
	Port int    `env:"HTTP_PORT"`
}

// HTTPClient - параметры подключения к HTTP через API Token
type HTTPClient struct {
	URL   string
	Token string
}

// DB - параметры подключения к базе данных
type DB struct {
	Driver   string `env:"DB_DRIVER"`
	Host     string `env:"DB_HOST"`
	Port     int    `env:"DB_PORT"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	Name     string `env:"DB_NAME"`
}

// Statsd - параметры сборщика метрик
type Statsd struct {
	Enabled bool   `env:"STATSD_ENABLED"`
	Host    string `env:"STATSD_HOST"`
	Port    int    `env:"STATSD_PORT"`
	Prefix  string `env:"STATSD_PREFIX"`
}

// LDAP - параметры подключения к LDAP
type LDAP struct {
	Enabled            bool   `env:"LDAP_ENABLED"`
	Host               string `env:"LDAP_HOST"`
	Port               int    `env:"LDAP_PORT"`
	BindDN             string `env:"LDAP_BINDDN"`
	BindPassword       string `env:"LDAP_PASSWORD"`
	BaseDN             string `env:"LDAP_BASEDN"`
	SkipTLS            bool   `env:"LDAP_SKIP_TLS"`
	UseSSL             bool   `env:"LDAP_USE_SSL"`
	InsecureSkipVerify bool   `env:"LDAP_SKIP_VERIFY"`
}

// DockerRegistry параметры подключения к реестру докера
type DockerRegistry struct {
	Username      string `env:"REGISTRY_USER"`
	Password      string `env:"REGISTRY_PASSWORD"`
	Auth          string `env:"REGISTRY_AUTH"`
	ServerAddress string `env:"REGISTRY_ADDRESS"`
	IdentityToken string `env:"IDENTITY_TOKEN"`
	RegistryToken string `env:"REGISTRY_TOKEN"`
}
