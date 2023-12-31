package domain

type ErrorStruct struct {
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

type MessageResponse struct {
	Message string `json:"message"`
}
