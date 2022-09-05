package utilities

import "os"

type ServerConfig struct {
	ServerAddress string
	ServerPort    string
	debug         bool
	AccessSecret  string
	RefreshSecret string
	MaxFileSize   uint64
	dbConfig      *databaseConfig
	cacheConfig   *cacheConfig
	keyConfig     *keyConfig
	tokenConfig   *tokenConfig
}

func NewServerConfig() *ServerConfig {
	serverConfig := ServerConfig{}
	if os.Getenv("oees_server") != "" || len(os.Getenv("oees_server")) != 0 {
		serverConfig.ServerAddress = os.Getenv("oees_server")
	} else {
		serverConfig.ServerAddress = "localhost"
	}
	if os.Getenv("oees_server_port") != "" || len(os.Getenv("oees_server_port")) != 0 {
		serverConfig.ServerPort = os.Getenv("oees_server_port")
	} else {
		serverConfig.ServerPort = "8080"
	}
	serverConfig.debug = false
	serverConfig.dbConfig = NewDatabaseConfig()
	serverConfig.cacheConfig = NewCacheConfig()
	serverConfig.keyConfig = NewKeyConfig()
	serverConfig.tokenConfig = NewTokenConfig()
	return &serverConfig
}

func (conf *ServerConfig) IsDebug() bool {
	return conf.debug
}

func (conf *ServerConfig) GetDatabaseConfig() *databaseConfig {
	return conf.dbConfig
}

func (conf *ServerConfig) GetCacheConfig() *cacheConfig {
	return conf.cacheConfig
}

func (conf *ServerConfig) GetKeyConfig() *keyConfig {
	return conf.keyConfig
}

func (conf *ServerConfig) GetTokenConfig() *tokenConfig {
	return conf.tokenConfig
}

type databaseConfig struct {
	DbHost            string
	DbPort            string
	DbName            string
	DbUser            string
	DbPassword        string
	WarehouseHost     string
	WarehouseDBName   string
	WarehouseUser     string
	WarehousePassword string
}

func NewDatabaseConfig() *databaseConfig {
	dbConf := databaseConfig{}
	if os.Getenv("oees_database_server") != "" || len(os.Getenv("oees_database_server")) != 0 {
		dbConf.DbHost = os.Getenv("oees_database_server")
	} else {
		dbConf.DbHost = "localhost"
	}
	if os.Getenv("oees_database_server_port") != "" || len(os.Getenv("oees_database_server_port")) != 0 {
		dbConf.DbPort = os.Getenv("oees_database_server_port")
	} else {
		dbConf.DbPort = "3306"
	}
	if os.Getenv("mysql_username") != "" || len(os.Getenv("mysql_username")) != 0 {
		dbConf.DbUser = os.Getenv("mysql_username")
	} else {
		dbConf.DbUser = ""
	}
	if os.Getenv("mysql_password") != "" || len(os.Getenv("mysql_password")) != 0 {
		dbConf.DbPassword = os.Getenv("mysql_password")
	} else {
		dbConf.DbPassword = ""
	}
	if os.Getenv("oees_database_name") != "" || len(os.Getenv("oees_database_name")) != 0 {
		dbConf.DbName = os.Getenv("oees_database_name")
	} else {
		dbConf.DbName = "oees"
	}
	if os.Getenv("warehouse_database_server") != "" || len(os.Getenv("warehouse_database_server")) != 0 {
		dbConf.WarehouseHost = os.Getenv("warehouse_database_server")
	} else {
		dbConf.WarehouseHost = "localhost"
	}
	if os.Getenv("warehouse_username") != "" || len(os.Getenv("warehouse_username")) != 0 {
		dbConf.WarehouseUser = os.Getenv("warehouse_username")
	} else {
		dbConf.WarehouseUser = ""
	}
	if os.Getenv("warehouse_password") != "" || len(os.Getenv("warehouse_password")) != 0 {
		dbConf.WarehousePassword = os.Getenv("warehouse_password")
	} else {
		dbConf.WarehousePassword = ""
	}
	if os.Getenv("warehouse_database_name") != "" || len(os.Getenv("warehouse_database_name")) != 0 {
		dbConf.WarehouseDBName = os.Getenv("warehouse_database_name")
	} else {
		dbConf.WarehouseDBName = ""
	}
	return &dbConf
}

type cacheConfig struct {
	CacheHost     string
	CachePort     string
	CachePassword string
}

func NewCacheConfig() *cacheConfig {
	cacheConf := cacheConfig{}
	if os.Getenv("cache_server") != "" || len(os.Getenv("cache_server")) != 0 {
		cacheConf.CacheHost = os.Getenv("cache_server")
	} else {
		cacheConf.CacheHost = "localhost"
	}
	if os.Getenv("cache_port") != "" || len(os.Getenv("cache_port")) != 0 {
		cacheConf.CachePort = os.Getenv("cache_port")
	} else {
		cacheConf.CachePort = "6379"
	}
	if os.Getenv("cache_password") != "" || len(os.Getenv("cache_password")) != 0 {
		cacheConf.CachePassword = os.Getenv("cache_password")
	} else {
		cacheConf.CachePassword = "rishabh2791"
	}
	return &cacheConf
}

type keyConfig struct {
	AccessTokenPrivateKeyPath  string
	AccessTokenPublicKeyPath   string
	RefreshTokenPrivateKeyPath string
	RefreshTokenPublicKeyPath  string
}

func NewKeyConfig() *keyConfig {
	return &keyConfig{
		AccessTokenPrivateKeyPath:  "/home/pi/Development/backend/access-private.pem",
		AccessTokenPublicKeyPath:   "/home/pi/Development/backend/access-public.pem",
		RefreshTokenPrivateKeyPath: "/home/pi/Development/backend/refresh-private.pem",
		RefreshTokenPublicKeyPath:  "/home/pi/Development/backend/refresh-public.pem",
	}
}

type tokenConfig struct {
	JWTExpiration     int // in minutes
	RefreshExpiration int // in days
}

func NewTokenConfig() *tokenConfig {
	return &tokenConfig{
		JWTExpiration:     10 * 2,
		RefreshExpiration: 7,
	}
}
