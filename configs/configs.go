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

type ValidatorConf struct {
	PasswordMin    int
	PasswordMax    int
	AdvertBodyMin  int
	AdvertBodyMax  int
	PriceMax       int
	AdvertTitleMax int
	AdvertTitleMin int
}

type EmailConf struct {
	Email    string
	Password string
	Smtp     string
	SmtpPort string
}
