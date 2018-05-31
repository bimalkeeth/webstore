package configuration

type Configuration struct{
	Provider string
	HostName string
	Port string
	User string
	Password string
	Database string
	Create bool
	Migrate bool
	ConnType int
}