package configs

type PostgresConf struct {
	DBName   string
	Username string
	Password string
	Host     string
}

type ServerConf struct {
	Addr      string
	RWTimeout int
	HeaderMBs int
}
