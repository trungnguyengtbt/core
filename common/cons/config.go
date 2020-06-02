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
		MongoUri      string
		MongoDatabase string
		MongoUser     string
		MongoPassword string
		MongoAuth     string
	}
)
