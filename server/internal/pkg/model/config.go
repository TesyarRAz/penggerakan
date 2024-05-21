package model

import "strconv"

type DotEnvConfig map[string]string

func (d DotEnvConfig) StringOrDefaultKey(key string, defKey string) string {
	val, ok := d[key]
	if !ok || val == "" {
		return d[defKey]
	}

	return d[key]
}

func (d DotEnvConfig) Modify(data map[string]string) DotEnvConfig {
	for k, v := range data {
		d[k] = v
	}

	return d
}

func (d DotEnvConfig) AppID() string {
	return d["APP_ID"]
}

func (d DotEnvConfig) JWTSecret() string {
	return d.StringOrDefaultKey("JWT_SECRET", "APP_ID")
}

func (d DotEnvConfig) JWTRefreshSecret() string {
	return d.StringOrDefaultKey("JWT_REFRESH_SECRET", "APP_ID")
}

func (d DotEnvConfig) DBHost() string {
	return d["DB_HOST"]
}

func (d DotEnvConfig) DBPort() int {
	port, err := strconv.Atoi(d["DB_PORT"])
	if err != nil {
		panic(err)
	}

	return port
}

func (d DotEnvConfig) DBUsername() string {
	return d["DB_USERNAME"]
}

func (d DotEnvConfig) DBPassword() string {
	return d["DB_PASSWORD"]
}

func (d DotEnvConfig) DBName() string {
	return d["DB_NAME"]
}

func (d DotEnvConfig) DBSSLMode() string {
	return d["DB_SSL_MODE"]
}

func (d DotEnvConfig) LogLevel() int {
	level, err := strconv.Atoi(d["LOG_LEVEL"])
	if err != nil {
		panic(err)
	}

	return level
}

func (d DotEnvConfig) WebPort() int {
	port, err := strconv.Atoi(d["WEB_PORT"])
	if err != nil {
		panic(err)
	}

	return port
}

func (d DotEnvConfig) RedisHost() string {
	return d["REDIS_HOST"]
}

func (d DotEnvConfig) RedisPort() int {
	port, err := strconv.Atoi(d["REDIS_PORT"])
	if err != nil {
		panic(err)
	}

	return port
}

func (d DotEnvConfig) RedisPassword() string {
	return d["REDIS_PASSWORD"]
}
