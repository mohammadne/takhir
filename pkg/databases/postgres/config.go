package postgres

type Config struct {
	Host     string `required:"true"`
	Port     int    `required:"true"`
	User     string `required:"true"`
	Password string `required:"true"`
	Database string `required:"true"`
}
