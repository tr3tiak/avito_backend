package entity

type Config struct {
	UserDB     string
	PasswordDB string
	NameDB     string
	Port       string
}

func NewConfig() *Config {
	Conf := Config{
		UserDB:     "Alexander",
		PasswordDB: "3u0wowLTSY",
		NameDB:     "ads",
		Port:       "8080",
	}
	return &Conf
}
