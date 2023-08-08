package domain

type ErrorStruct struct {
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

type SignInResponse struct {
	Data struct {
		Token string `json:"token"`
	}
}

type MessageResponse struct {
	Data struct {
		Message string `json:"message"`
	}
}
