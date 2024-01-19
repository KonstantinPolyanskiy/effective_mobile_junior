package handler

const (
	savePersonType    = "save person"
	getPersonType     = "get person"
	deletePersonType  = "delete person"
	updatePersonType  = "update person"
	marshallingType   = "marshalling"
	convertType       = "convert"
	stringToIntAction = "input is not a string"
)

type AppError struct {
	Type   string `json:"type"`
	Action string `json:"action"`
}
