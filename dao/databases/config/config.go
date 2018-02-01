package config

//DBConfig is Common processing of the DB
type DBConfig interface {
	DSN() string
}

//Config is base DB Config struct
type Config struct {
	Host     string
	Port     int
	User     string
	Password string
}
