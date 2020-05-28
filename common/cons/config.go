package cons

type (
	DbConfiguration struct {
		DbDriverName string
		DbHost       string
		DbPort       string
		DbUsername   string
		DbPassword   string
		DbName       string
	}

	MongoConfiguration struct {
		MgUri        string
		MgDB         string
		MgUserName   string
		MgPassword   string
		MgAuthSource string
	}
)
