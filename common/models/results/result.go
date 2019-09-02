package results

type value struct {
	Code     int
	Messages string
}

type list struct {
	Success               value
	InternalError         value
	BadRequest            value
	ExistData             value
	NotFoundData	value
	InvalidAuthentication value
}

var Results = &list{
	Success:               value{90, "Success"},
	InternalError:         value{91, "Unknown"},
	BadRequest:            value{92, "Bad request format"},
	ExistData:             value{93, "Data existed"},
	NotFoundData: value{94,"Data not found"},
	InvalidAuthentication: value{95, "Invalid Authentication"},
}
