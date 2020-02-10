package config

// SQLDatabase to store Sql database config
type SQLDatabase struct {
	Name       string
	User       string
	Password   string
	Port       string
	Connection string
	Host       string
}

type MongoDatabase struct {
	Connection string
	Database   string
}
