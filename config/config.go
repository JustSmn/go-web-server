/*

	Application settings

*/

package config

type Config struct {
	ServerAddress string
	DatabaseURL   string
}

func Load() Config {
	return Config{
		ServerAddress: ":8080",
		DatabaseURL:   "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable",
	}
}
