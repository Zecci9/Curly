package config

type Config struct {
	AppName string
	Port    string
}

func Load() Config {
	return Config{
		AppName: "Curly",
		Port:    "8080",
	}
}