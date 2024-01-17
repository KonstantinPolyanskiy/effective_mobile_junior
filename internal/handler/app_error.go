package handler

type AppError struct {
	Type   string `json:"type"`
	Action string `json:"action"`
}
