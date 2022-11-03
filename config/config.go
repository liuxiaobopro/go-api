package config

type Config struct {
	App      App
	Database Database
	Upload   Upload
	Ws       Ws
}

var Conf = new(Config)
